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

package plc4go

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	options "github.com/apache/plc4x/plc4go/spi/options"

	transports "github.com/apache/plc4x/plc4go/spi/transports"

	url "net/url"
)

// MockPlcDriver is an autogenerated mock type for the PlcDriver type
type MockPlcDriver struct {
	mock.Mock
}

type MockPlcDriver_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcDriver) EXPECT() *MockPlcDriver_Expecter {
	return &MockPlcDriver_Expecter{mock: &_m.Mock}
}

// CheckQuery provides a mock function with given fields: query
func (_m *MockPlcDriver) CheckQuery(query string) error {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for CheckQuery")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcDriver_CheckQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckQuery'
type MockPlcDriver_CheckQuery_Call struct {
	*mock.Call
}

// CheckQuery is a helper method to define mock.On call
//   - query string
func (_e *MockPlcDriver_Expecter) CheckQuery(query interface{}) *MockPlcDriver_CheckQuery_Call {
	return &MockPlcDriver_CheckQuery_Call{Call: _e.mock.On("CheckQuery", query)}
}

func (_c *MockPlcDriver_CheckQuery_Call) Run(run func(query string)) *MockPlcDriver_CheckQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcDriver_CheckQuery_Call) Return(_a0 error) *MockPlcDriver_CheckQuery_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriver_CheckQuery_Call) RunAndReturn(run func(string) error) *MockPlcDriver_CheckQuery_Call {
	_c.Call.Return(run)
	return _c
}

// CheckTagAddress provides a mock function with given fields: tagAddress
func (_m *MockPlcDriver) CheckTagAddress(tagAddress string) error {
	ret := _m.Called(tagAddress)

	if len(ret) == 0 {
		panic("no return value specified for CheckTagAddress")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(tagAddress)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcDriver_CheckTagAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckTagAddress'
type MockPlcDriver_CheckTagAddress_Call struct {
	*mock.Call
}

// CheckTagAddress is a helper method to define mock.On call
//   - tagAddress string
func (_e *MockPlcDriver_Expecter) CheckTagAddress(tagAddress interface{}) *MockPlcDriver_CheckTagAddress_Call {
	return &MockPlcDriver_CheckTagAddress_Call{Call: _e.mock.On("CheckTagAddress", tagAddress)}
}

func (_c *MockPlcDriver_CheckTagAddress_Call) Run(run func(tagAddress string)) *MockPlcDriver_CheckTagAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcDriver_CheckTagAddress_Call) Return(_a0 error) *MockPlcDriver_CheckTagAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriver_CheckTagAddress_Call) RunAndReturn(run func(string) error) *MockPlcDriver_CheckTagAddress_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockPlcDriver) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcDriver_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockPlcDriver_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockPlcDriver_Expecter) Close() *MockPlcDriver_Close_Call {
	return &MockPlcDriver_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockPlcDriver_Close_Call) Run(run func()) *MockPlcDriver_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcDriver_Close_Call) Return(_a0 error) *MockPlcDriver_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriver_Close_Call) RunAndReturn(run func() error) *MockPlcDriver_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Discover provides a mock function with given fields: callback, discoveryOptions
func (_m *MockPlcDriver) Discover(callback func(model.PlcDiscoveryItem), discoveryOptions ...options.WithDiscoveryOption) error {
	_va := make([]interface{}, len(discoveryOptions))
	for _i := range discoveryOptions {
		_va[_i] = discoveryOptions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, callback)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Discover")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(func(model.PlcDiscoveryItem), ...options.WithDiscoveryOption) error); ok {
		r0 = rf(callback, discoveryOptions...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcDriver_Discover_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Discover'
type MockPlcDriver_Discover_Call struct {
	*mock.Call
}

// Discover is a helper method to define mock.On call
//   - callback func(model.PlcDiscoveryItem)
//   - discoveryOptions ...options.WithDiscoveryOption
func (_e *MockPlcDriver_Expecter) Discover(callback interface{}, discoveryOptions ...interface{}) *MockPlcDriver_Discover_Call {
	return &MockPlcDriver_Discover_Call{Call: _e.mock.On("Discover",
		append([]interface{}{callback}, discoveryOptions...)...)}
}

func (_c *MockPlcDriver_Discover_Call) Run(run func(callback func(model.PlcDiscoveryItem), discoveryOptions ...options.WithDiscoveryOption)) *MockPlcDriver_Discover_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]options.WithDiscoveryOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(options.WithDiscoveryOption)
			}
		}
		run(args[0].(func(model.PlcDiscoveryItem)), variadicArgs...)
	})
	return _c
}

func (_c *MockPlcDriver_Discover_Call) Return(_a0 error) *MockPlcDriver_Discover_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriver_Discover_Call) RunAndReturn(run func(func(model.PlcDiscoveryItem), ...options.WithDiscoveryOption) error) *MockPlcDriver_Discover_Call {
	_c.Call.Return(run)
	return _c
}

// DiscoverWithContext provides a mock function with given fields: ctx, callback, discoveryOptions
func (_m *MockPlcDriver) DiscoverWithContext(ctx context.Context, callback func(model.PlcDiscoveryItem), discoveryOptions ...options.WithDiscoveryOption) error {
	_va := make([]interface{}, len(discoveryOptions))
	for _i := range discoveryOptions {
		_va[_i] = discoveryOptions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, callback)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DiscoverWithContext")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(model.PlcDiscoveryItem), ...options.WithDiscoveryOption) error); ok {
		r0 = rf(ctx, callback, discoveryOptions...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcDriver_DiscoverWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DiscoverWithContext'
type MockPlcDriver_DiscoverWithContext_Call struct {
	*mock.Call
}

// DiscoverWithContext is a helper method to define mock.On call
//   - ctx context.Context
//   - callback func(model.PlcDiscoveryItem)
//   - discoveryOptions ...options.WithDiscoveryOption
func (_e *MockPlcDriver_Expecter) DiscoverWithContext(ctx interface{}, callback interface{}, discoveryOptions ...interface{}) *MockPlcDriver_DiscoverWithContext_Call {
	return &MockPlcDriver_DiscoverWithContext_Call{Call: _e.mock.On("DiscoverWithContext",
		append([]interface{}{ctx, callback}, discoveryOptions...)...)}
}

func (_c *MockPlcDriver_DiscoverWithContext_Call) Run(run func(ctx context.Context, callback func(model.PlcDiscoveryItem), discoveryOptions ...options.WithDiscoveryOption)) *MockPlcDriver_DiscoverWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]options.WithDiscoveryOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(options.WithDiscoveryOption)
			}
		}
		run(args[0].(context.Context), args[1].(func(model.PlcDiscoveryItem)), variadicArgs...)
	})
	return _c
}

func (_c *MockPlcDriver_DiscoverWithContext_Call) Return(_a0 error) *MockPlcDriver_DiscoverWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriver_DiscoverWithContext_Call) RunAndReturn(run func(context.Context, func(model.PlcDiscoveryItem), ...options.WithDiscoveryOption) error) *MockPlcDriver_DiscoverWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnection provides a mock function with given fields: transportUrl, _a1, _a2
func (_m *MockPlcDriver) GetConnection(transportUrl url.URL, _a1 map[string]transports.Transport, _a2 map[string][]string) <-chan PlcConnectionConnectResult {
	ret := _m.Called(transportUrl, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for GetConnection")
	}

	var r0 <-chan PlcConnectionConnectResult
	if rf, ok := ret.Get(0).(func(url.URL, map[string]transports.Transport, map[string][]string) <-chan PlcConnectionConnectResult); ok {
		r0 = rf(transportUrl, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan PlcConnectionConnectResult)
		}
	}

	return r0
}

// MockPlcDriver_GetConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnection'
type MockPlcDriver_GetConnection_Call struct {
	*mock.Call
}

// GetConnection is a helper method to define mock.On call
//   - transportUrl url.URL
//   - _a1 map[string]transports.Transport
//   - _a2 map[string][]string
func (_e *MockPlcDriver_Expecter) GetConnection(transportUrl interface{}, _a1 interface{}, _a2 interface{}) *MockPlcDriver_GetConnection_Call {
	return &MockPlcDriver_GetConnection_Call{Call: _e.mock.On("GetConnection", transportUrl, _a1, _a2)}
}

func (_c *MockPlcDriver_GetConnection_Call) Run(run func(transportUrl url.URL, _a1 map[string]transports.Transport, _a2 map[string][]string)) *MockPlcDriver_GetConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(url.URL), args[1].(map[string]transports.Transport), args[2].(map[string][]string))
	})
	return _c
}

func (_c *MockPlcDriver_GetConnection_Call) Return(_a0 <-chan PlcConnectionConnectResult) *MockPlcDriver_GetConnection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriver_GetConnection_Call) RunAndReturn(run func(url.URL, map[string]transports.Transport, map[string][]string) <-chan PlcConnectionConnectResult) *MockPlcDriver_GetConnection_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnectionWithContext provides a mock function with given fields: ctx, transportUrl, _a2, driverOptions
func (_m *MockPlcDriver) GetConnectionWithContext(ctx context.Context, transportUrl url.URL, _a2 map[string]transports.Transport, driverOptions map[string][]string) <-chan PlcConnectionConnectResult {
	ret := _m.Called(ctx, transportUrl, _a2, driverOptions)

	if len(ret) == 0 {
		panic("no return value specified for GetConnectionWithContext")
	}

	var r0 <-chan PlcConnectionConnectResult
	if rf, ok := ret.Get(0).(func(context.Context, url.URL, map[string]transports.Transport, map[string][]string) <-chan PlcConnectionConnectResult); ok {
		r0 = rf(ctx, transportUrl, _a2, driverOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan PlcConnectionConnectResult)
		}
	}

	return r0
}

// MockPlcDriver_GetConnectionWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnectionWithContext'
type MockPlcDriver_GetConnectionWithContext_Call struct {
	*mock.Call
}

// GetConnectionWithContext is a helper method to define mock.On call
//   - ctx context.Context
//   - transportUrl url.URL
//   - _a2 map[string]transports.Transport
//   - driverOptions map[string][]string
func (_e *MockPlcDriver_Expecter) GetConnectionWithContext(ctx interface{}, transportUrl interface{}, _a2 interface{}, driverOptions interface{}) *MockPlcDriver_GetConnectionWithContext_Call {
	return &MockPlcDriver_GetConnectionWithContext_Call{Call: _e.mock.On("GetConnectionWithContext", ctx, transportUrl, _a2, driverOptions)}
}

func (_c *MockPlcDriver_GetConnectionWithContext_Call) Run(run func(ctx context.Context, transportUrl url.URL, _a2 map[string]transports.Transport, driverOptions map[string][]string)) *MockPlcDriver_GetConnectionWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(url.URL), args[2].(map[string]transports.Transport), args[3].(map[string][]string))
	})
	return _c
}

func (_c *MockPlcDriver_GetConnectionWithContext_Call) Return(_a0 <-chan PlcConnectionConnectResult) *MockPlcDriver_GetConnectionWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriver_GetConnectionWithContext_Call) RunAndReturn(run func(context.Context, url.URL, map[string]transports.Transport, map[string][]string) <-chan PlcConnectionConnectResult) *MockPlcDriver_GetConnectionWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefaultTransport provides a mock function with given fields:
func (_m *MockPlcDriver) GetDefaultTransport() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDefaultTransport")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcDriver_GetDefaultTransport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefaultTransport'
type MockPlcDriver_GetDefaultTransport_Call struct {
	*mock.Call
}

// GetDefaultTransport is a helper method to define mock.On call
func (_e *MockPlcDriver_Expecter) GetDefaultTransport() *MockPlcDriver_GetDefaultTransport_Call {
	return &MockPlcDriver_GetDefaultTransport_Call{Call: _e.mock.On("GetDefaultTransport")}
}

func (_c *MockPlcDriver_GetDefaultTransport_Call) Run(run func()) *MockPlcDriver_GetDefaultTransport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcDriver_GetDefaultTransport_Call) Return(_a0 string) *MockPlcDriver_GetDefaultTransport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriver_GetDefaultTransport_Call) RunAndReturn(run func() string) *MockPlcDriver_GetDefaultTransport_Call {
	_c.Call.Return(run)
	return _c
}

// GetProtocolCode provides a mock function with given fields:
func (_m *MockPlcDriver) GetProtocolCode() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetProtocolCode")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcDriver_GetProtocolCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProtocolCode'
type MockPlcDriver_GetProtocolCode_Call struct {
	*mock.Call
}

// GetProtocolCode is a helper method to define mock.On call
func (_e *MockPlcDriver_Expecter) GetProtocolCode() *MockPlcDriver_GetProtocolCode_Call {
	return &MockPlcDriver_GetProtocolCode_Call{Call: _e.mock.On("GetProtocolCode")}
}

func (_c *MockPlcDriver_GetProtocolCode_Call) Run(run func()) *MockPlcDriver_GetProtocolCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcDriver_GetProtocolCode_Call) Return(_a0 string) *MockPlcDriver_GetProtocolCode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriver_GetProtocolCode_Call) RunAndReturn(run func() string) *MockPlcDriver_GetProtocolCode_Call {
	_c.Call.Return(run)
	return _c
}

// GetProtocolName provides a mock function with given fields:
func (_m *MockPlcDriver) GetProtocolName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetProtocolName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcDriver_GetProtocolName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProtocolName'
type MockPlcDriver_GetProtocolName_Call struct {
	*mock.Call
}

// GetProtocolName is a helper method to define mock.On call
func (_e *MockPlcDriver_Expecter) GetProtocolName() *MockPlcDriver_GetProtocolName_Call {
	return &MockPlcDriver_GetProtocolName_Call{Call: _e.mock.On("GetProtocolName")}
}

func (_c *MockPlcDriver_GetProtocolName_Call) Run(run func()) *MockPlcDriver_GetProtocolName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcDriver_GetProtocolName_Call) Return(_a0 string) *MockPlcDriver_GetProtocolName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriver_GetProtocolName_Call) RunAndReturn(run func() string) *MockPlcDriver_GetProtocolName_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcDriver) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcDriver_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcDriver_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcDriver_Expecter) String() *MockPlcDriver_String_Call {
	return &MockPlcDriver_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcDriver_String_Call) Run(run func()) *MockPlcDriver_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcDriver_String_Call) Return(_a0 string) *MockPlcDriver_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriver_String_Call) RunAndReturn(run func() string) *MockPlcDriver_String_Call {
	_c.Call.Return(run)
	return _c
}

// SupportsDiscovery provides a mock function with given fields:
func (_m *MockPlcDriver) SupportsDiscovery() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SupportsDiscovery")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcDriver_SupportsDiscovery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SupportsDiscovery'
type MockPlcDriver_SupportsDiscovery_Call struct {
	*mock.Call
}

// SupportsDiscovery is a helper method to define mock.On call
func (_e *MockPlcDriver_Expecter) SupportsDiscovery() *MockPlcDriver_SupportsDiscovery_Call {
	return &MockPlcDriver_SupportsDiscovery_Call{Call: _e.mock.On("SupportsDiscovery")}
}

func (_c *MockPlcDriver_SupportsDiscovery_Call) Run(run func()) *MockPlcDriver_SupportsDiscovery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcDriver_SupportsDiscovery_Call) Return(_a0 bool) *MockPlcDriver_SupportsDiscovery_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcDriver_SupportsDiscovery_Call) RunAndReturn(run func() bool) *MockPlcDriver_SupportsDiscovery_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcDriver creates a new instance of MockPlcDriver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcDriver(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcDriver {
	mock := &MockPlcDriver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
