// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/containernetworking/cni/pkg/types (interfaces: Result)

package mock_types

import (
	types "github.com/containernetworking/cni/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Result interface
type MockResult struct {
	ctrl     *gomock.Controller
	recorder *_MockResultRecorder
}

// Recorder for MockResult (not exported)
type _MockResultRecorder struct {
	mock *MockResult
}

func NewMockResult(ctrl *gomock.Controller) *MockResult {
	mock := &MockResult{ctrl: ctrl}
	mock.recorder = &_MockResultRecorder{mock}
	return mock
}

func (_m *MockResult) EXPECT() *_MockResultRecorder {
	return _m.recorder
}

func (_m *MockResult) GetAsVersion(_param0 string) (types.Result, error) {
	ret := _m.ctrl.Call(_m, "GetAsVersion", _param0)
	ret0, _ := ret[0].(types.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockResultRecorder) GetAsVersion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAsVersion", arg0)
}

func (_m *MockResult) Print() error {
	ret := _m.ctrl.Call(_m, "Print")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockResultRecorder) Print() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Print")
}

func (_m *MockResult) String() string {
	ret := _m.ctrl.Call(_m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockResultRecorder) String() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "String")
}

func (_m *MockResult) Version() string {
	ret := _m.ctrl.Call(_m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockResultRecorder) Version() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Version")
}
