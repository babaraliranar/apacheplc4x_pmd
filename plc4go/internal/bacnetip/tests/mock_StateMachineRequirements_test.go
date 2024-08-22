/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.42.2. DO NOT EDIT.

package tests

import (
	bacnetip "github.com/apache/plc4x/plc4go/internal/bacnetip"
	mock "github.com/stretchr/testify/mock"
)

// MockStateMachineRequirements is an autogenerated mock type for the StateMachineRequirements type
type MockStateMachineRequirements struct {
	mock.Mock
}

type MockStateMachineRequirements_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStateMachineRequirements) EXPECT() *MockStateMachineRequirements_Expecter {
	return &MockStateMachineRequirements_Expecter{mock: &_m.Mock}
}

// Send provides a mock function with given fields: args, kwargs
func (_m *MockStateMachineRequirements) Send(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	ret := _m.Called(args, kwargs)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(bacnetip.Args, bacnetip.KWArgs) error); ok {
		r0 = rf(args, kwargs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStateMachineRequirements_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockStateMachineRequirements_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - args bacnetip.Args
//   - kwargs bacnetip.KWArgs
func (_e *MockStateMachineRequirements_Expecter) Send(args interface{}, kwargs interface{}) *MockStateMachineRequirements_Send_Call {
	return &MockStateMachineRequirements_Send_Call{Call: _e.mock.On("Send", args, kwargs)}
}

func (_c *MockStateMachineRequirements_Send_Call) Run(run func(args bacnetip.Args, kwargs bacnetip.KWArgs)) *MockStateMachineRequirements_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bacnetip.Args), args[1].(bacnetip.KWArgs))
	})
	return _c
}

func (_c *MockStateMachineRequirements_Send_Call) Return(_a0 error) *MockStateMachineRequirements_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStateMachineRequirements_Send_Call) RunAndReturn(run func(bacnetip.Args, bacnetip.KWArgs) error) *MockStateMachineRequirements_Send_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStateMachineRequirements creates a new instance of MockStateMachineRequirements. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStateMachineRequirements(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStateMachineRequirements {
	mock := &MockStateMachineRequirements{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
