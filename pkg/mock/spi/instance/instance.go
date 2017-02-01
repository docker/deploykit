// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/docker/infrakit/pkg/spi/instance (interfaces: Plugin)

package instance

import (
	instance "github.com/docker/infrakit/pkg/spi/instance"
	types "github.com/docker/infrakit/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Plugin interface
type MockPlugin struct {
	ctrl     *gomock.Controller
	recorder *_MockPluginRecorder
}

// Recorder for MockPlugin (not exported)
type _MockPluginRecorder struct {
	mock *MockPlugin
}

func NewMockPlugin(ctrl *gomock.Controller) *MockPlugin {
	mock := &MockPlugin{ctrl: ctrl}
	mock.recorder = &_MockPluginRecorder{mock}
	return mock
}

func (_m *MockPlugin) EXPECT() *_MockPluginRecorder {
	return _m.recorder
}

func (_m *MockPlugin) DescribeInstances(_param0 map[string]string) ([]instance.Description, error) {
	ret := _m.ctrl.Call(_m, "DescribeInstances", _param0)
	ret0, _ := ret[0].([]instance.Description)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPluginRecorder) DescribeInstances(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeInstances", arg0)
}

func (_m *MockPlugin) Destroy(_param0 instance.ID) error {
	ret := _m.ctrl.Call(_m, "Destroy", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockPluginRecorder) Destroy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Destroy", arg0)
}

func (_m *MockPlugin) Label(_param0 instance.ID, _param1 map[string]string) error {
	ret := _m.ctrl.Call(_m, "Label", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockPluginRecorder) Label(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Label", arg0, arg1)
}

func (_m *MockPlugin) Provision(_param0 instance.Spec) (*instance.ID, error) {
	ret := _m.ctrl.Call(_m, "Provision", _param0)
	ret0, _ := ret[0].(*instance.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPluginRecorder) Provision(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Provision", arg0)
}

func (_m *MockPlugin) Validate(_param0 *types.Any) error {
	ret := _m.ctrl.Call(_m, "Validate", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockPluginRecorder) Validate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Validate", arg0)
}
