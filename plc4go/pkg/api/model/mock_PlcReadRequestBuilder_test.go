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

// MockPlcReadRequestBuilder is an autogenerated mock type for the PlcReadRequestBuilder type
type MockPlcReadRequestBuilder struct {
	mock.Mock
}

type MockPlcReadRequestBuilder_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcReadRequestBuilder) EXPECT() *MockPlcReadRequestBuilder_Expecter {
	return &MockPlcReadRequestBuilder_Expecter{mock: &_m.Mock}
}

// AddTag provides a mock function with given fields: tagName, tag
func (_m *MockPlcReadRequestBuilder) AddTag(tagName string, tag PlcTag) PlcReadRequestBuilder {
	ret := _m.Called(tagName, tag)

	var r0 PlcReadRequestBuilder
	if rf, ok := ret.Get(0).(func(string, PlcTag) PlcReadRequestBuilder); ok {
		r0 = rf(tagName, tag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcReadRequestBuilder)
		}
	}

	return r0
}

// MockPlcReadRequestBuilder_AddTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTag'
type MockPlcReadRequestBuilder_AddTag_Call struct {
	*mock.Call
}

// AddTag is a helper method to define mock.On call
//   - tagName string
//   - tag PlcTag
func (_e *MockPlcReadRequestBuilder_Expecter) AddTag(tagName interface{}, tag interface{}) *MockPlcReadRequestBuilder_AddTag_Call {
	return &MockPlcReadRequestBuilder_AddTag_Call{Call: _e.mock.On("AddTag", tagName, tag)}
}

func (_c *MockPlcReadRequestBuilder_AddTag_Call) Run(run func(tagName string, tag PlcTag)) *MockPlcReadRequestBuilder_AddTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(PlcTag))
	})
	return _c
}

func (_c *MockPlcReadRequestBuilder_AddTag_Call) Return(_a0 PlcReadRequestBuilder) *MockPlcReadRequestBuilder_AddTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadRequestBuilder_AddTag_Call) RunAndReturn(run func(string, PlcTag) PlcReadRequestBuilder) *MockPlcReadRequestBuilder_AddTag_Call {
	_c.Call.Return(run)
	return _c
}

// AddTagAddress provides a mock function with given fields: tagName, tagAddress
func (_m *MockPlcReadRequestBuilder) AddTagAddress(tagName string, tagAddress string) PlcReadRequestBuilder {
	ret := _m.Called(tagName, tagAddress)

	var r0 PlcReadRequestBuilder
	if rf, ok := ret.Get(0).(func(string, string) PlcReadRequestBuilder); ok {
		r0 = rf(tagName, tagAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcReadRequestBuilder)
		}
	}

	return r0
}

// MockPlcReadRequestBuilder_AddTagAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTagAddress'
type MockPlcReadRequestBuilder_AddTagAddress_Call struct {
	*mock.Call
}

// AddTagAddress is a helper method to define mock.On call
//   - tagName string
//   - tagAddress string
func (_e *MockPlcReadRequestBuilder_Expecter) AddTagAddress(tagName interface{}, tagAddress interface{}) *MockPlcReadRequestBuilder_AddTagAddress_Call {
	return &MockPlcReadRequestBuilder_AddTagAddress_Call{Call: _e.mock.On("AddTagAddress", tagName, tagAddress)}
}

func (_c *MockPlcReadRequestBuilder_AddTagAddress_Call) Run(run func(tagName string, tagAddress string)) *MockPlcReadRequestBuilder_AddTagAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockPlcReadRequestBuilder_AddTagAddress_Call) Return(_a0 PlcReadRequestBuilder) *MockPlcReadRequestBuilder_AddTagAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcReadRequestBuilder_AddTagAddress_Call) RunAndReturn(run func(string, string) PlcReadRequestBuilder) *MockPlcReadRequestBuilder_AddTagAddress_Call {
	_c.Call.Return(run)
	return _c
}

// Build provides a mock function with given fields:
func (_m *MockPlcReadRequestBuilder) Build() (PlcReadRequest, error) {
	ret := _m.Called()

	var r0 PlcReadRequest
	var r1 error
	if rf, ok := ret.Get(0).(func() (PlcReadRequest, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() PlcReadRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcReadRequest)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPlcReadRequestBuilder_Build_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Build'
type MockPlcReadRequestBuilder_Build_Call struct {
	*mock.Call
}

// Build is a helper method to define mock.On call
func (_e *MockPlcReadRequestBuilder_Expecter) Build() *MockPlcReadRequestBuilder_Build_Call {
	return &MockPlcReadRequestBuilder_Build_Call{Call: _e.mock.On("Build")}
}

func (_c *MockPlcReadRequestBuilder_Build_Call) Run(run func()) *MockPlcReadRequestBuilder_Build_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcReadRequestBuilder_Build_Call) Return(_a0 PlcReadRequest, _a1 error) *MockPlcReadRequestBuilder_Build_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPlcReadRequestBuilder_Build_Call) RunAndReturn(run func() (PlcReadRequest, error)) *MockPlcReadRequestBuilder_Build_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPlcReadRequestBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPlcReadRequestBuilder creates a new instance of MockPlcReadRequestBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPlcReadRequestBuilder(t mockConstructorTestingTNewMockPlcReadRequestBuilder) *MockPlcReadRequestBuilder {
	mock := &MockPlcReadRequestBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
