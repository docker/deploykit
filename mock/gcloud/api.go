// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/docker/infrakit.gcp/plugin/instance/gcloud (interfaces: GCloud)

package gcloud

import (
	gcloud "github.com/docker/infrakit.gcp/plugin/instance/gcloud"
	gomock "github.com/golang/mock/gomock"
	v1 "google.golang.org/api/compute/v1"
)

// Mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockGCloudRecorder
}

// Recorder for MockGCloud (not exported)
type _MockGCloudRecorder struct {
	mock *MockAPI
}

func NewMockGCloud(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &_MockGCloudRecorder{mock}
	return mock
}

func (_m *MockAPI) EXPECT() *_MockGCloudRecorder {
	return _m.recorder
}

func (_m *MockAPI) AddInstanceToTargetPool(_param0 string, _param1 ...string) error {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "AddInstanceToTargetPool", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockGCloudRecorder) AddInstanceToTargetPool(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddInstanceToTargetPool", _s...)
}

func (_m *MockAPI) CreateInstance(_param0 string, _param1 *gcloud.InstanceSettings) error {
	ret := _m.ctrl.Call(_m, "CreateInstance", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockGCloudRecorder) CreateInstance(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateInstance", arg0, arg1)
}

func (_m *MockAPI) DeleteInstance(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteInstance", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockGCloudRecorder) DeleteInstance(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteInstance", arg0)
}

func (_m *MockAPI) ListInstances() ([]*v1.Instance, error) {
	ret := _m.ctrl.Call(_m, "ListInstances")
	ret0, _ := ret[0].([]*v1.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockGCloudRecorder) ListInstances() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListInstances")
}
