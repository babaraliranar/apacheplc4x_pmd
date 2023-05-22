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

// Code generated by mockery v2.27.1. DO NOT EDIT.

package model

import mock "github.com/stretchr/testify/mock"

// MockPlcConsumerRegistration is an autogenerated mock type for the PlcConsumerRegistration type
type MockPlcConsumerRegistration struct {
	mock.Mock
}

type MockPlcConsumerRegistration_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcConsumerRegistration) EXPECT() *MockPlcConsumerRegistration_Expecter {
	return &MockPlcConsumerRegistration_Expecter{mock: &_m.Mock}
}

// GetConsumerId provides a mock function with given fields:
func (_m *MockPlcConsumerRegistration) GetConsumerId() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockPlcConsumerRegistration_GetConsumerId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConsumerId'
type MockPlcConsumerRegistration_GetConsumerId_Call struct {
	*mock.Call
}

// GetConsumerId is a helper method to define mock.On call
func (_e *MockPlcConsumerRegistration_Expecter) GetConsumerId() *MockPlcConsumerRegistration_GetConsumerId_Call {
	return &MockPlcConsumerRegistration_GetConsumerId_Call{Call: _e.mock.On("GetConsumerId")}
}

func (_c *MockPlcConsumerRegistration_GetConsumerId_Call) Run(run func()) *MockPlcConsumerRegistration_GetConsumerId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConsumerRegistration_GetConsumerId_Call) Return(_a0 int) *MockPlcConsumerRegistration_GetConsumerId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConsumerRegistration_GetConsumerId_Call) RunAndReturn(run func() int) *MockPlcConsumerRegistration_GetConsumerId_Call {
	_c.Call.Return(run)
	return _c
}

// GetSubscriptionHandles provides a mock function with given fields:
func (_m *MockPlcConsumerRegistration) GetSubscriptionHandles() []PlcSubscriptionHandle {
	ret := _m.Called()

	var r0 []PlcSubscriptionHandle
	if rf, ok := ret.Get(0).(func() []PlcSubscriptionHandle); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]PlcSubscriptionHandle)
		}
	}

	return r0
}

// MockPlcConsumerRegistration_GetSubscriptionHandles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubscriptionHandles'
type MockPlcConsumerRegistration_GetSubscriptionHandles_Call struct {
	*mock.Call
}

// GetSubscriptionHandles is a helper method to define mock.On call
func (_e *MockPlcConsumerRegistration_Expecter) GetSubscriptionHandles() *MockPlcConsumerRegistration_GetSubscriptionHandles_Call {
	return &MockPlcConsumerRegistration_GetSubscriptionHandles_Call{Call: _e.mock.On("GetSubscriptionHandles")}
}

func (_c *MockPlcConsumerRegistration_GetSubscriptionHandles_Call) Run(run func()) *MockPlcConsumerRegistration_GetSubscriptionHandles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConsumerRegistration_GetSubscriptionHandles_Call) Return(_a0 []PlcSubscriptionHandle) *MockPlcConsumerRegistration_GetSubscriptionHandles_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConsumerRegistration_GetSubscriptionHandles_Call) RunAndReturn(run func() []PlcSubscriptionHandle) *MockPlcConsumerRegistration_GetSubscriptionHandles_Call {
	_c.Call.Return(run)
	return _c
}

// Unregister provides a mock function with given fields:
func (_m *MockPlcConsumerRegistration) Unregister() {
	_m.Called()
}

// MockPlcConsumerRegistration_Unregister_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unregister'
type MockPlcConsumerRegistration_Unregister_Call struct {
	*mock.Call
}

// Unregister is a helper method to define mock.On call
func (_e *MockPlcConsumerRegistration_Expecter) Unregister() *MockPlcConsumerRegistration_Unregister_Call {
	return &MockPlcConsumerRegistration_Unregister_Call{Call: _e.mock.On("Unregister")}
}

func (_c *MockPlcConsumerRegistration_Unregister_Call) Run(run func()) *MockPlcConsumerRegistration_Unregister_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConsumerRegistration_Unregister_Call) Return() *MockPlcConsumerRegistration_Unregister_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPlcConsumerRegistration_Unregister_Call) RunAndReturn(run func()) *MockPlcConsumerRegistration_Unregister_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPlcConsumerRegistration interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPlcConsumerRegistration creates a new instance of MockPlcConsumerRegistration. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPlcConsumerRegistration(t mockConstructorTestingTNewMockPlcConsumerRegistration) *MockPlcConsumerRegistration {
	mock := &MockPlcConsumerRegistration{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
