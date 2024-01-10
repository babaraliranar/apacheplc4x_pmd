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

package tracer

import mock "github.com/stretchr/testify/mock"

// MockProvider is an autogenerated mock type for the Provider type
type MockProvider struct {
	mock.Mock
}

type MockProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProvider) EXPECT() *MockProvider_Expecter {
	return &MockProvider_Expecter{mock: &_m.Mock}
}

// EnableTracer provides a mock function with given fields:
func (_m *MockProvider) EnableTracer() {
	_m.Called()
}

// MockProvider_EnableTracer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnableTracer'
type MockProvider_EnableTracer_Call struct {
	*mock.Call
}

// EnableTracer is a helper method to define mock.On call
func (_e *MockProvider_Expecter) EnableTracer() *MockProvider_EnableTracer_Call {
	return &MockProvider_EnableTracer_Call{Call: _e.mock.On("EnableTracer")}
}

func (_c *MockProvider_EnableTracer_Call) Run(run func()) *MockProvider_EnableTracer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProvider_EnableTracer_Call) Return() *MockProvider_EnableTracer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockProvider_EnableTracer_Call) RunAndReturn(run func()) *MockProvider_EnableTracer_Call {
	_c.Call.Return(run)
	return _c
}

// GetTracer provides a mock function with given fields:
func (_m *MockProvider) GetTracer() *Tracer {
	ret := _m.Called()

	var r0 *Tracer
	if rf, ok := ret.Get(0).(func() *Tracer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Tracer)
		}
	}

	return r0
}

// MockProvider_GetTracer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTracer'
type MockProvider_GetTracer_Call struct {
	*mock.Call
}

// GetTracer is a helper method to define mock.On call
func (_e *MockProvider_Expecter) GetTracer() *MockProvider_GetTracer_Call {
	return &MockProvider_GetTracer_Call{Call: _e.mock.On("GetTracer")}
}

func (_c *MockProvider_GetTracer_Call) Run(run func()) *MockProvider_GetTracer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProvider_GetTracer_Call) Return(_a0 *Tracer) *MockProvider_GetTracer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockProvider_GetTracer_Call) RunAndReturn(run func() *Tracer) *MockProvider_GetTracer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProvider creates a new instance of MockProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProvider {
	mock := &MockProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
