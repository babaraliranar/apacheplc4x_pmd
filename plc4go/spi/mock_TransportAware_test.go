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

// Code generated by mockery v2.32.4. DO NOT EDIT.

package spi

import (
	transports "github.com/apache/plc4x/plc4go/spi/transports"
	mock "github.com/stretchr/testify/mock"
)

// MockTransportAware is an autogenerated mock type for the TransportAware type
type MockTransportAware struct {
	mock.Mock
}

type MockTransportAware_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTransportAware) EXPECT() *MockTransportAware_Expecter {
	return &MockTransportAware_Expecter{mock: &_m.Mock}
}

// GetTransport provides a mock function with given fields: transportName, connectionString, options
func (_m *MockTransportAware) GetTransport(transportName string, connectionString string, options map[string][]string) (transports.Transport, error) {
	ret := _m.Called(transportName, connectionString, options)

	var r0 transports.Transport
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, map[string][]string) (transports.Transport, error)); ok {
		return rf(transportName, connectionString, options)
	}
	if rf, ok := ret.Get(0).(func(string, string, map[string][]string) transports.Transport); ok {
		r0 = rf(transportName, connectionString, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(transports.Transport)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, map[string][]string) error); ok {
		r1 = rf(transportName, connectionString, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransportAware_GetTransport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransport'
type MockTransportAware_GetTransport_Call struct {
	*mock.Call
}

// GetTransport is a helper method to define mock.On call
//   - transportName string
//   - connectionString string
//   - options map[string][]string
func (_e *MockTransportAware_Expecter) GetTransport(transportName interface{}, connectionString interface{}, options interface{}) *MockTransportAware_GetTransport_Call {
	return &MockTransportAware_GetTransport_Call{Call: _e.mock.On("GetTransport", transportName, connectionString, options)}
}

func (_c *MockTransportAware_GetTransport_Call) Run(run func(transportName string, connectionString string, options map[string][]string)) *MockTransportAware_GetTransport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(map[string][]string))
	})
	return _c
}

func (_c *MockTransportAware_GetTransport_Call) Return(_a0 transports.Transport, _a1 error) *MockTransportAware_GetTransport_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransportAware_GetTransport_Call) RunAndReturn(run func(string, string, map[string][]string) (transports.Transport, error)) *MockTransportAware_GetTransport_Call {
	_c.Call.Return(run)
	return _c
}

// ListTransportNames provides a mock function with given fields:
func (_m *MockTransportAware) ListTransportNames() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockTransportAware_ListTransportNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTransportNames'
type MockTransportAware_ListTransportNames_Call struct {
	*mock.Call
}

// ListTransportNames is a helper method to define mock.On call
func (_e *MockTransportAware_Expecter) ListTransportNames() *MockTransportAware_ListTransportNames_Call {
	return &MockTransportAware_ListTransportNames_Call{Call: _e.mock.On("ListTransportNames")}
}

func (_c *MockTransportAware_ListTransportNames_Call) Run(run func()) *MockTransportAware_ListTransportNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTransportAware_ListTransportNames_Call) Return(_a0 []string) *MockTransportAware_ListTransportNames_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTransportAware_ListTransportNames_Call) RunAndReturn(run func() []string) *MockTransportAware_ListTransportNames_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterTransport provides a mock function with given fields: transport
func (_m *MockTransportAware) RegisterTransport(transport transports.Transport) {
	_m.Called(transport)
}

// MockTransportAware_RegisterTransport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterTransport'
type MockTransportAware_RegisterTransport_Call struct {
	*mock.Call
}

// RegisterTransport is a helper method to define mock.On call
//   - transport transports.Transport
func (_e *MockTransportAware_Expecter) RegisterTransport(transport interface{}) *MockTransportAware_RegisterTransport_Call {
	return &MockTransportAware_RegisterTransport_Call{Call: _e.mock.On("RegisterTransport", transport)}
}

func (_c *MockTransportAware_RegisterTransport_Call) Run(run func(transport transports.Transport)) *MockTransportAware_RegisterTransport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(transports.Transport))
	})
	return _c
}

func (_c *MockTransportAware_RegisterTransport_Call) Return() *MockTransportAware_RegisterTransport_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTransportAware_RegisterTransport_Call) RunAndReturn(run func(transports.Transport)) *MockTransportAware_RegisterTransport_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTransportAware creates a new instance of MockTransportAware. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTransportAware(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTransportAware {
	mock := &MockTransportAware{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
