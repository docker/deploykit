package swarm

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/client"
	docker_types "github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/libmachete/plugin/group/types"
	"github.com/docker/libmachete/plugin/group/util"
	"github.com/docker/libmachete/spi/group"
	"github.com/docker/libmachete/spi/instance"
	"golang.org/x/net/context"
	"text/template"
)

const (
	roleWorker  = "worker"
	roleManager = "manager"
)

// NewSwarmProvisionHelper creates a ProvisionHelper that creates manager and worker nodes connected in a swarm.
func NewSwarmProvisionHelper() types.ProvisionHelper {
	return &swarmProvisioner{}
}

type swarmProvisioner struct {
	dockerClient client.Client
}

// TODO(wfarner): Tag instances with a UUID, and tag the Docker engine with the same UUID.  We will use this to
// associate swarm nodes with instances.

// TODO(wfarner): Add a ProvisionHelper function to check the health of an instance.  Use the Swarm node association
// (see TODO above) to provide this.

func (s swarmProvisioner) Validate(config group.Configuration, parsed types.Schema) error {
	if config.Role == roleManager {
		if len(parsed.IPs) != 1 && len(parsed.IPs) != 3 && len(parsed.IPs) != 5 {
			return errors.New("Must have 1, 3, or 5 managers")
		}
	}
	return nil
}

func (s swarmProvisioner) GroupKind(roleName string) (types.GroupKind, error) {
	switch roleName {
	case roleWorker:
		return types.KindDynamicIP, nil
	case roleManager:
		return types.KindStaticIP, nil
	default:
		return types.KindNone, errors.New("Unsupported role type")
	}
}

const (
	// associationTag is a machine tag added to associate machines with Swarm nodes.
	associationTag = "swarm-association-id"

	// bootScript is used to generate node boot scripts.
	bootScript = fmt.Sprintf(`#!/bin/sh
set -o errexit
set -o nounset
set -o xtrace

mkdir -p /etc/docker
cat << EOF > /etc/docker/daemon.json
{
  "labels": ["%s={{.ASSOCIATION_ID}}"],
}
EOF

docker swarm join {{.MY_IP}} --token {{.JOIN_TOKEN}}
`, associationTag)
)

func generateBootScript(joinIP, joinToken, associationID string) string {
	buffer := bytes.Buffer{}
	templ := template.Must(template.New("").Parse(bootScript))
	err := templ.Execute(&buffer, map[string]string{
		"MY_IP":          joinIP,
		"JOIN_TOKEN":     joinToken,
		"ASSOCIATION_ID": associationID,
	})
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

// Healthy determines whether an instance is healthy.  This is determined by whether it has successfully joined the
// Swarm.
func (s swarmProvisioner) Healthy(inst instance.Description) (bool, error) {
	associationID, exists := inst.Tags[associationTag]
	if !exists {
		log.Info("Reporting unhealthy for instance without an association tag", inst.ID)
		return false, nil
	}

	filter := filters.NewArgs()
	filter.Add("label", fmt.Sprintf("%s=%s", associationTag, associationID))

	nodes, err := s.dockerClient.NodeList(context.Background(), docker_types.NodeListOptions{Filter: filter})
	if err != nil {
		return false, err
	}

	if len(nodes) > 1 {
		log.Warnf("Expected at most one node with label %s, but found %s", associationID, nodes)
	}

	return len(nodes) == 1, nil
}

func (s swarmProvisioner) PreProvision(
	config group.Configuration,
	details types.ProvisionDetails) (types.ProvisionDetails, error) {

	ctx := context.Background()
	swarmStatus, err := s.dockerClient.SwarmInspect(ctx)
	if err != nil {
		return details, fmt.Errorf("Failed to fetch Swarm join tokens: %s", err)
	}

	self, _, err := s.dockerClient.NodeInspectWithRaw(ctx, "self")
	if err != nil {
		return details, fmt.Errorf("Failed to fetch Swarm node status: %s", err)
	}

	if self.ManagerStatus == nil {
		return details, errors.New(
			"Swarm node status did not include manager status.  Need to run 'docker swarm init`?")
	}

	associationID := util.RandomAlphaNumericString(8)
	details.Tags[associationTag] = associationID

	switch config.Role {
	case roleWorker:
		details.BootScript = generateBootScript(
			self.ManagerStatus.Addr,
			swarmStatus.JoinTokens.Worker,
			associationID)

	case roleManager:
		if details.PrivateIP == nil {
			return details, errors.New("Manager nodes require an assigned private IP address")
		}

		details.BootScript = generateBootScript(
			self.ManagerStatus.Addr,
			swarmStatus.JoinTokens.Manager,
			associationID)

		volume := instance.VolumeID(*details.PrivateIP)
		details.Volume = &volume

	default:
		return details, errors.New("Unsupported role type")
	}

	// TODO(wfarner): Tag with with the Swarm cluster UUID for scoping.
	details.Tags["swarm-id"] = swarmStatus.ID

	return details, nil
}
