package group

import (
	"encoding/json"
	"net/http"

	"github.com/docker/infrakit/pkg/plugin"
	"github.com/docker/infrakit/pkg/spi/group"
)

// PluginServer returns a RPCService that conforms to the net/rpc rpc call convention.
func PluginServer(p group.Plugin) *Group {
	return &Group{plugin: p}
}

// Group the exported type needed to conform to json-rpc call convention
type Group struct {
	plugin group.Plugin
}

// Info returns a metadata object about the plugin, if the plugin implements it.  See plugin.Vendor
func (p *Group) Info() plugin.Info {
	if m, is := p.plugin.(plugin.Vendor); is {
		return m.Info()
	}
	return plugin.NoInfo
}

// ExampleProperties returns an example properties used by the plugin
func (p *Group) ExampleProperties() *json.RawMessage {
	if i, is := p.plugin.(plugin.InputExample); is {
		return i.ExampleProperties()
	}
	return nil
}

// CommitGroup is the rpc method to commit a group
func (p *Group) CommitGroup(_ *http.Request, req *CommitGroupRequest, resp *CommitGroupResponse) error {
	details, err := p.plugin.CommitGroup(req.Spec, req.Pretend)
	if err != nil {
		return err
	}
	resp.Details = details
	return nil
}

// FreeGroup is the rpc method to free a group
func (p *Group) FreeGroup(_ *http.Request, req *FreeGroupRequest, resp *FreeGroupResponse) error {
	err := p.plugin.FreeGroup(req.ID)
	if err != nil {
		return err
	}
	resp.OK = true
	return nil
}

// DescribeGroup is the rpc method to describe a group
func (p *Group) DescribeGroup(_ *http.Request, req *DescribeGroupRequest, resp *DescribeGroupResponse) error {
	desc, err := p.plugin.DescribeGroup(req.ID)
	if err != nil {
		return err
	}
	resp.Description = desc
	return nil
}

// DestroyGroup is the rpc method to destroy a group
func (p *Group) DestroyGroup(_ *http.Request, req *DestroyGroupRequest, resp *DestroyGroupResponse) error {
	err := p.plugin.DestroyGroup(req.ID)
	if err != nil {
		return err
	}
	resp.OK = true
	return nil
}

// InspectGroups is the rpc method to inspect groups
func (p *Group) InspectGroups(_ *http.Request, req *InspectGroupsRequest, resp *InspectGroupsResponse) error {
	groups, err := p.plugin.InspectGroups()
	if err != nil {
		return err
	}
	resp.Groups = groups
	return nil
}