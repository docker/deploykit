package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	docker_types "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/docker/infrakit/pkg/spi/flavor"
	"github.com/docker/infrakit/pkg/spi/instance"
	"github.com/docker/infrakit/pkg/template"
	"github.com/docker/infrakit/pkg/types"
	"golang.org/x/net/context"
)

const (
	ebsAttachment string = "ebs"
)

// Spec is the value passed in the `Properties` field of configs
type Spec struct {

	// Attachments indicate the devices that are to be attached to the instance
	Attachments map[instance.LogicalID][]instance.Attachment

	// InitScriptTemplateURL overrides the template specified when the plugin started up.
	InitScriptTemplateURL string
}

func validateIDsAndAttachments(logicalIDs []instance.LogicalID,
	attachments map[instance.LogicalID][]instance.Attachment) error {

	// Each attachment association must be represented by a logical ID.
	idsMap := map[instance.LogicalID]bool{}
	for _, id := range logicalIDs {
		if _, exists := idsMap[id]; exists {
			return fmt.Errorf("LogicalID %v specified more than once", id)
		}

		idsMap[id] = true
	}
	for id := range attachments {
		if _, exists := idsMap[id]; !exists {
			return fmt.Errorf("LogicalID %v used for an attachment but is not in group LogicalIDs", id)
		}
	}

	// Only EBS attachments are supported.
	for _, atts := range attachments {
		for _, attachment := range atts {
			if attachment.Type == "" {
				return fmt.Errorf(
					"Attachment Type %s must be specified for '%s'",
					ebsAttachment,
					attachment.ID)
			}

			if attachment.Type != ebsAttachment {
				return fmt.Errorf(
					"Invalid attachment Type '%s', only %s is supported",
					attachment.Type,
					ebsAttachment)
			}
		}
	}

	// Each attachment may only be used once.
	allAttachmentIDs := map[string]bool{}
	for _, atts := range attachments {
		for _, attachment := range atts {
			if _, exists := allAttachmentIDs[attachment.ID]; exists {
				return fmt.Errorf("Attachment %v specified more than once", attachment.ID)
			}
			allAttachmentIDs[attachment.ID] = true
		}
	}

	return nil
}

func swarmState(docker client.APIClient) (status swarm.Swarm, node swarm.Node, err error) {
	ctx := context.Background()
	info, err := docker.Info(ctx)
	if err != nil {
		return
	}
	node, _, err = docker.NodeInspectWithRaw(ctx, info.Swarm.NodeID)
	if err != nil {
		return
	}
	status, err = docker.SwarmInspect(ctx)
	return
}

func exportTemplateFunctions(swarmStatus swarm.Swarm, nodeInfo swarm.Node, link types.Link) []template.Function {

	// Get a single consistent view of the data across multiple calls by exporting functions that
	// query the input state

	return []template.Function{
		{
			Name:        "INFRAKIT_LABELS",
			Description: "The label name to use for linking an InfraKit managed resource somewhere else.",
			Func: func() []string {
				return link.KVPairs()
			},
		},
		{
			Name:        "SWARM_MANAGER_IP",
			Description: "The label name to use for linking an InfraKit managed resource somewhere else.",
			Func: func() (string, error) {
				if nodeInfo.ManagerStatus == nil {
					return "", fmt.Errorf("no manager status")
				}
				return nodeInfo.ManagerStatus.Addr, nil
			},
		},
		{
			Name:        "SWARM_INITIALIZED",
			Description: "Returns true if the swarm has been initialized.",
			Func: func() bool {
				return nodeInfo.ManagerStatus != nil
			},
		},
		{
			Name:        "SWARM_JOIN_TOKENS",
			Description: "Returns the swarm JoinTokens object, with either .Manager or .Worker fields",
			Func: func() interface{} {
				return swarmStatus.JoinTokens
			},
		},
		{
			Name:        "SWARM_CLUSTER_ID",
			Description: "Returns the swarm cluster UUID",
			Func: func() interface{} {
				return swarmStatus.ID
			},
		},
	}
}

// Healthy determines whether an instance is healthy.  This is determined by whether it has successfully joined the
// Swarm.
func healthy(client client.APIClient, inst instance.Description) (flavor.Health, error) {

	link := types.NewLinkFromMap(inst.Tags)
	if !link.Valid() {
		log.Info("Reporting unhealthy for instance without an association tag", inst.ID)
		return flavor.Unhealthy, nil
	}

	filter := filters.NewArgs()
	filter.Add("label", fmt.Sprintf("%s=%s", link.Label(), link.Value()))

	nodes, err := client.NodeList(context.Background(), docker_types.NodeListOptions{Filters: filter})
	if err != nil {
		return flavor.Unknown, err
	}

	switch {
	case len(nodes) == 0:
		// The instance may not yet be joined, so we consider the health unknown.
		return flavor.Unknown, nil

	case len(nodes) == 1:
		return flavor.Healthy, nil

	default:
		log.Warnf("Expected at most one node with label %s, but found %s", link.Value(), nodes)
		return flavor.Healthy, nil
	}
}
