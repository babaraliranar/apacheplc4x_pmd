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

package bacnetip

import (
	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"

	values "github.com/apache/plc4x/plc4go/pkg/api/values"
)

// MockBacNetPlcTag is an autogenerated mock type for the BacNetPlcTag type
type MockBacNetPlcTag struct {
	mock.Mock
}

type MockBacNetPlcTag_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBacNetPlcTag) EXPECT() *MockBacNetPlcTag_Expecter {
	return &MockBacNetPlcTag_Expecter{mock: &_m.Mock}
}

// GetAddressString provides a mock function with given fields:
func (_m *MockBacNetPlcTag) GetAddressString() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockBacNetPlcTag_GetAddressString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAddressString'
type MockBacNetPlcTag_GetAddressString_Call struct {
	*mock.Call
}

// GetAddressString is a helper method to define mock.On call
func (_e *MockBacNetPlcTag_Expecter) GetAddressString() *MockBacNetPlcTag_GetAddressString_Call {
	return &MockBacNetPlcTag_GetAddressString_Call{Call: _e.mock.On("GetAddressString")}
}

func (_c *MockBacNetPlcTag_GetAddressString_Call) Run(run func()) *MockBacNetPlcTag_GetAddressString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBacNetPlcTag_GetAddressString_Call) Return(_a0 string) *MockBacNetPlcTag_GetAddressString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBacNetPlcTag_GetAddressString_Call) RunAndReturn(run func() string) *MockBacNetPlcTag_GetAddressString_Call {
	_c.Call.Return(run)
	return _c
}

// GetArrayInfo provides a mock function with given fields:
func (_m *MockBacNetPlcTag) GetArrayInfo() []model.ArrayInfo {
	ret := _m.Called()

	var r0 []model.ArrayInfo
	if rf, ok := ret.Get(0).(func() []model.ArrayInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ArrayInfo)
		}
	}

	return r0
}

// MockBacNetPlcTag_GetArrayInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArrayInfo'
type MockBacNetPlcTag_GetArrayInfo_Call struct {
	*mock.Call
}

// GetArrayInfo is a helper method to define mock.On call
func (_e *MockBacNetPlcTag_Expecter) GetArrayInfo() *MockBacNetPlcTag_GetArrayInfo_Call {
	return &MockBacNetPlcTag_GetArrayInfo_Call{Call: _e.mock.On("GetArrayInfo")}
}

func (_c *MockBacNetPlcTag_GetArrayInfo_Call) Run(run func()) *MockBacNetPlcTag_GetArrayInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBacNetPlcTag_GetArrayInfo_Call) Return(_a0 []model.ArrayInfo) *MockBacNetPlcTag_GetArrayInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBacNetPlcTag_GetArrayInfo_Call) RunAndReturn(run func() []model.ArrayInfo) *MockBacNetPlcTag_GetArrayInfo_Call {
	_c.Call.Return(run)
	return _c
}

// GetObjectId provides a mock function with given fields:
func (_m *MockBacNetPlcTag) GetObjectId() objectId {
	ret := _m.Called()

	var r0 objectId
	if rf, ok := ret.Get(0).(func() objectId); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(objectId)
	}

	return r0
}

// MockBacNetPlcTag_GetObjectId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObjectId'
type MockBacNetPlcTag_GetObjectId_Call struct {
	*mock.Call
}

// GetObjectId is a helper method to define mock.On call
func (_e *MockBacNetPlcTag_Expecter) GetObjectId() *MockBacNetPlcTag_GetObjectId_Call {
	return &MockBacNetPlcTag_GetObjectId_Call{Call: _e.mock.On("GetObjectId")}
}

func (_c *MockBacNetPlcTag_GetObjectId_Call) Run(run func()) *MockBacNetPlcTag_GetObjectId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBacNetPlcTag_GetObjectId_Call) Return(_a0 objectId) *MockBacNetPlcTag_GetObjectId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBacNetPlcTag_GetObjectId_Call) RunAndReturn(run func() objectId) *MockBacNetPlcTag_GetObjectId_Call {
	_c.Call.Return(run)
	return _c
}

// GetProperties provides a mock function with given fields:
func (_m *MockBacNetPlcTag) GetProperties() []property {
	ret := _m.Called()

	var r0 []property
	if rf, ok := ret.Get(0).(func() []property); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]property)
		}
	}

	return r0
}

// MockBacNetPlcTag_GetProperties_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProperties'
type MockBacNetPlcTag_GetProperties_Call struct {
	*mock.Call
}

// GetProperties is a helper method to define mock.On call
func (_e *MockBacNetPlcTag_Expecter) GetProperties() *MockBacNetPlcTag_GetProperties_Call {
	return &MockBacNetPlcTag_GetProperties_Call{Call: _e.mock.On("GetProperties")}
}

func (_c *MockBacNetPlcTag_GetProperties_Call) Run(run func()) *MockBacNetPlcTag_GetProperties_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBacNetPlcTag_GetProperties_Call) Return(_a0 []property) *MockBacNetPlcTag_GetProperties_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBacNetPlcTag_GetProperties_Call) RunAndReturn(run func() []property) *MockBacNetPlcTag_GetProperties_Call {
	_c.Call.Return(run)
	return _c
}

// GetValueType provides a mock function with given fields:
func (_m *MockBacNetPlcTag) GetValueType() values.PlcValueType {
	ret := _m.Called()

	var r0 values.PlcValueType
	if rf, ok := ret.Get(0).(func() values.PlcValueType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(values.PlcValueType)
	}

	return r0
}

// MockBacNetPlcTag_GetValueType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValueType'
type MockBacNetPlcTag_GetValueType_Call struct {
	*mock.Call
}

// GetValueType is a helper method to define mock.On call
func (_e *MockBacNetPlcTag_Expecter) GetValueType() *MockBacNetPlcTag_GetValueType_Call {
	return &MockBacNetPlcTag_GetValueType_Call{Call: _e.mock.On("GetValueType")}
}

func (_c *MockBacNetPlcTag_GetValueType_Call) Run(run func()) *MockBacNetPlcTag_GetValueType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBacNetPlcTag_GetValueType_Call) Return(_a0 values.PlcValueType) *MockBacNetPlcTag_GetValueType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBacNetPlcTag_GetValueType_Call) RunAndReturn(run func() values.PlcValueType) *MockBacNetPlcTag_GetValueType_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockBacNetPlcTag interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockBacNetPlcTag creates a new instance of MockBacNetPlcTag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockBacNetPlcTag(t mockConstructorTestingTNewMockBacNetPlcTag) *MockBacNetPlcTag {
	mock := &MockBacNetPlcTag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
