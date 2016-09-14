package swarm

import (
	"fmt"
	docker_types "github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/swarm"
	mock_client "github.com/docker/libmachete/mock/docker/engine-api/client"
	"github.com/docker/libmachete/plugin/group/types"
	"github.com/docker/libmachete/spi/group"
	"github.com/docker/libmachete/spi/instance"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGroupKind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	helper := NewSwarmProvisionHelper(mock_client.NewMockAPIClient(ctrl))

	require.Equal(t, types.KindDynamicIP, helper.GroupKind("worker"))
	require.Equal(t, types.KindStaticIP, helper.GroupKind("manager"))
	require.Equal(t, types.KindUnknown, helper.GroupKind("other"))
}

func TestAssociation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_client.NewMockAPIClient(ctrl)

	helper := NewSwarmProvisionHelper(client)

	swarmInfo := swarm.Swarm{
		ClusterInfo: swarm.ClusterInfo{ID: "ClusterUUID"},
		JoinTokens: swarm.JoinTokens{
			Manager: "ManagerToken",
			Worker:  "WorkerToken",
		},
	}
	client.EXPECT().SwarmInspect(gomock.Any()).Return(swarmInfo, nil)
	nodeInfo := swarm.Node{ManagerStatus: &swarm.ManagerStatus{Addr: "1.2.3.4"}}
	client.EXPECT().NodeInspectWithRaw(gomock.Any(), "self").Return(nodeInfo, nil, nil)

	details, err := helper.PreProvision(
		group.Configuration{Role: "worker"},
		types.ProvisionDetails{Tags: map[string]string{"a": "b"}})
	require.NoError(t, err)
	require.Equal(t, "b", details.Tags["a"])
	associationID := details.Tags[associationTag]
	require.NotEqual(t, "", associationID)

	// Perform a rudimentary check to ensure that the expected fields are in the BootScript, without having any
	// other knowledge about the script structure.
	require.Contains(t, details.BootScript, associationID)
	require.Contains(t, details.BootScript, swarmInfo.JoinTokens.Worker)
	require.Contains(t, details.BootScript, nodeInfo.ManagerStatus.Addr)

	// An instance with no association information is considered unhealthy.
	healthy, err := helper.Healthy(instance.Description{})
	require.NoError(t, err)
	require.False(t, healthy)

	filter, err := filters.FromParam(fmt.Sprintf(`{"label": {"%s=%s": true}}`, associationTag, associationID))
	require.NoError(t, err)
	client.EXPECT().NodeList(gomock.Any(), docker_types.NodeListOptions{Filter: filter}).Return(
		[]swarm.Node{
			{},
		}, nil)
	healthy, err = helper.Healthy(instance.Description{Tags: map[string]string{associationTag: associationID}})
	require.NoError(t, err)
	require.True(t, healthy)
}
