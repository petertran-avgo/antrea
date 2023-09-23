// Copyright 2023 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: antrea.io/antrea/pkg/agent/memberlist (interfaces: Memberlist)
//
// Generated by this command:
//
//	mockgen -copyright_file hack/boilerplate/license_header.raw.txt -destination pkg/agent/memberlist/mock_memberlist_test.go -package memberlist antrea.io/antrea/pkg/agent/memberlist Memberlist
//
// Package memberlist is a generated GoMock package.
package memberlist

import (
	reflect "reflect"
	time "time"

	memberlist "github.com/hashicorp/memberlist"
	gomock "go.uber.org/mock/gomock"
)

// MockMemberlist is a mock of Memberlist interface.
type MockMemberlist struct {
	ctrl     *gomock.Controller
	recorder *MockMemberlistMockRecorder
}

// MockMemberlistMockRecorder is the mock recorder for MockMemberlist.
type MockMemberlistMockRecorder struct {
	mock *MockMemberlist
}

// NewMockMemberlist creates a new mock instance.
func NewMockMemberlist(ctrl *gomock.Controller) *MockMemberlist {
	mock := &MockMemberlist{ctrl: ctrl}
	mock.recorder = &MockMemberlistMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemberlist) EXPECT() *MockMemberlistMockRecorder {
	return m.recorder
}

// Join mocks base method.
func (m *MockMemberlist) Join(arg0 []string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Join", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Join indicates an expected call of Join.
func (mr *MockMemberlistMockRecorder) Join(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockMemberlist)(nil).Join), arg0)
}

// Leave mocks base method.
func (m *MockMemberlist) Leave(arg0 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leave", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Leave indicates an expected call of Leave.
func (mr *MockMemberlistMockRecorder) Leave(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leave", reflect.TypeOf((*MockMemberlist)(nil).Leave), arg0)
}

// Members mocks base method.
func (m *MockMemberlist) Members() []*memberlist.Node {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Members")
	ret0, _ := ret[0].([]*memberlist.Node)
	return ret0
}

// Members indicates an expected call of Members.
func (mr *MockMemberlistMockRecorder) Members() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Members", reflect.TypeOf((*MockMemberlist)(nil).Members))
}

// Shutdown mocks base method.
func (m *MockMemberlist) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockMemberlistMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockMemberlist)(nil).Shutdown))
}
