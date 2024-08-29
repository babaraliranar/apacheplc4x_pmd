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

package bacnetip

import mock "github.com/stretchr/testify/mock"

// MockIOQControllerRequirements is an autogenerated mock type for the IOQControllerRequirements type
type MockIOQControllerRequirements struct {
	mock.Mock
}

type MockIOQControllerRequirements_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIOQControllerRequirements) EXPECT() *MockIOQControllerRequirements_Expecter {
	return &MockIOQControllerRequirements_Expecter{mock: &_m.Mock}
}

// Abort provides a mock function with given fields: err
func (_m *MockIOQControllerRequirements) Abort(err error) error {
	ret := _m.Called(err)

	if len(ret) == 0 {
		panic("no return value specified for Abort")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(error) error); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIOQControllerRequirements_Abort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Abort'
type MockIOQControllerRequirements_Abort_Call struct {
	*mock.Call
}

// Abort is a helper method to define mock.On call
//   - err error
func (_e *MockIOQControllerRequirements_Expecter) Abort(err interface{}) *MockIOQControllerRequirements_Abort_Call {
	return &MockIOQControllerRequirements_Abort_Call{Call: _e.mock.On("Abort", err)}
}

func (_c *MockIOQControllerRequirements_Abort_Call) Run(run func(err error)) *MockIOQControllerRequirements_Abort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *MockIOQControllerRequirements_Abort_Call) Return(_a0 error) *MockIOQControllerRequirements_Abort_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIOQControllerRequirements_Abort_Call) RunAndReturn(run func(error) error) *MockIOQControllerRequirements_Abort_Call {
	_c.Call.Return(run)
	return _c
}

// AbortIO provides a mock function with given fields: iocb, err
func (_m *MockIOQControllerRequirements) AbortIO(iocb _IOCB, err error) error {
	ret := _m.Called(iocb, err)

	if len(ret) == 0 {
		panic("no return value specified for AbortIO")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(_IOCB, error) error); ok {
		r0 = rf(iocb, err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIOQControllerRequirements_AbortIO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AbortIO'
type MockIOQControllerRequirements_AbortIO_Call struct {
	*mock.Call
}

// AbortIO is a helper method to define mock.On call
//   - iocb _IOCB
//   - err error
func (_e *MockIOQControllerRequirements_Expecter) AbortIO(iocb interface{}, err interface{}) *MockIOQControllerRequirements_AbortIO_Call {
	return &MockIOQControllerRequirements_AbortIO_Call{Call: _e.mock.On("AbortIO", iocb, err)}
}

func (_c *MockIOQControllerRequirements_AbortIO_Call) Run(run func(iocb _IOCB, err error)) *MockIOQControllerRequirements_AbortIO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_IOCB), args[1].(error))
	})
	return _c
}

func (_c *MockIOQControllerRequirements_AbortIO_Call) Return(_a0 error) *MockIOQControllerRequirements_AbortIO_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIOQControllerRequirements_AbortIO_Call) RunAndReturn(run func(_IOCB, error) error) *MockIOQControllerRequirements_AbortIO_Call {
	_c.Call.Return(run)
	return _c
}

// CompleteIO provides a mock function with given fields: iocb, pdu
func (_m *MockIOQControllerRequirements) CompleteIO(iocb _IOCB, pdu PDU) error {
	ret := _m.Called(iocb, pdu)

	if len(ret) == 0 {
		panic("no return value specified for CompleteIO")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(_IOCB, PDU) error); ok {
		r0 = rf(iocb, pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIOQControllerRequirements_CompleteIO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompleteIO'
type MockIOQControllerRequirements_CompleteIO_Call struct {
	*mock.Call
}

// CompleteIO is a helper method to define mock.On call
//   - iocb _IOCB
//   - pdu PDU
func (_e *MockIOQControllerRequirements_Expecter) CompleteIO(iocb interface{}, pdu interface{}) *MockIOQControllerRequirements_CompleteIO_Call {
	return &MockIOQControllerRequirements_CompleteIO_Call{Call: _e.mock.On("CompleteIO", iocb, pdu)}
}

func (_c *MockIOQControllerRequirements_CompleteIO_Call) Run(run func(iocb _IOCB, pdu PDU)) *MockIOQControllerRequirements_CompleteIO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_IOCB), args[1].(PDU))
	})
	return _c
}

func (_c *MockIOQControllerRequirements_CompleteIO_Call) Return(_a0 error) *MockIOQControllerRequirements_CompleteIO_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIOQControllerRequirements_CompleteIO_Call) RunAndReturn(run func(_IOCB, PDU) error) *MockIOQControllerRequirements_CompleteIO_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessIO provides a mock function with given fields: iocb
func (_m *MockIOQControllerRequirements) ProcessIO(iocb _IOCB) error {
	ret := _m.Called(iocb)

	if len(ret) == 0 {
		panic("no return value specified for ProcessIO")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(_IOCB) error); ok {
		r0 = rf(iocb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIOQControllerRequirements_ProcessIO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessIO'
type MockIOQControllerRequirements_ProcessIO_Call struct {
	*mock.Call
}

// ProcessIO is a helper method to define mock.On call
//   - iocb _IOCB
func (_e *MockIOQControllerRequirements_Expecter) ProcessIO(iocb interface{}) *MockIOQControllerRequirements_ProcessIO_Call {
	return &MockIOQControllerRequirements_ProcessIO_Call{Call: _e.mock.On("ProcessIO", iocb)}
}

func (_c *MockIOQControllerRequirements_ProcessIO_Call) Run(run func(iocb _IOCB)) *MockIOQControllerRequirements_ProcessIO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_IOCB))
	})
	return _c
}

func (_c *MockIOQControllerRequirements_ProcessIO_Call) Return(_a0 error) *MockIOQControllerRequirements_ProcessIO_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIOQControllerRequirements_ProcessIO_Call) RunAndReturn(run func(_IOCB) error) *MockIOQControllerRequirements_ProcessIO_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIOQControllerRequirements creates a new instance of MockIOQControllerRequirements. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIOQControllerRequirements(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIOQControllerRequirements {
	mock := &MockIOQControllerRequirements{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
