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
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockCompletableFuture is an autogenerated mock type for the CompletableFuture type
type MockCompletableFuture[T interface{}] struct {
	mock.Mock
}

type MockCompletableFuture_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *MockCompletableFuture[T]) EXPECT() *MockCompletableFuture_Expecter[T] {
	return &MockCompletableFuture_Expecter[T]{mock: &_m.Mock}
}

// Cancel provides a mock function with given fields:
func (_m *MockCompletableFuture[T]) Cancel() {
	_m.Called()
}

// MockCompletableFuture_Cancel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Cancel'
type MockCompletableFuture_Cancel_Call[T interface{}] struct {
	*mock.Call
}

// Cancel is a helper method to define mock.On call
func (_e *MockCompletableFuture_Expecter[T]) Cancel() *MockCompletableFuture_Cancel_Call[T] {
	return &MockCompletableFuture_Cancel_Call[T]{Call: _e.mock.On("Cancel")}
}

func (_c *MockCompletableFuture_Cancel_Call[T]) Run(run func()) *MockCompletableFuture_Cancel_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompletableFuture_Cancel_Call[T]) Return() *MockCompletableFuture_Cancel_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *MockCompletableFuture_Cancel_Call[T]) RunAndReturn(run func()) *MockCompletableFuture_Cancel_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Complete provides a mock function with given fields: value
func (_m *MockCompletableFuture[T]) Complete(value T) {
	_m.Called(value)
}

// MockCompletableFuture_Complete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Complete'
type MockCompletableFuture_Complete_Call[T interface{}] struct {
	*mock.Call
}

// Complete is a helper method to define mock.On call
//   - value T
func (_e *MockCompletableFuture_Expecter[T]) Complete(value interface{}) *MockCompletableFuture_Complete_Call[T] {
	return &MockCompletableFuture_Complete_Call[T]{Call: _e.mock.On("Complete", value)}
}

func (_c *MockCompletableFuture_Complete_Call[T]) Run(run func(value T)) *MockCompletableFuture_Complete_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(T))
	})
	return _c
}

func (_c *MockCompletableFuture_Complete_Call[T]) Return() *MockCompletableFuture_Complete_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *MockCompletableFuture_Complete_Call[T]) RunAndReturn(run func(T)) *MockCompletableFuture_Complete_Call[T] {
	_c.Call.Return(run)
	return _c
}

// CompleteWithError provides a mock function with given fields: err
func (_m *MockCompletableFuture[T]) CompleteWithError(err error) {
	_m.Called(err)
}

// MockCompletableFuture_CompleteWithError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompleteWithError'
type MockCompletableFuture_CompleteWithError_Call[T interface{}] struct {
	*mock.Call
}

// CompleteWithError is a helper method to define mock.On call
//   - err error
func (_e *MockCompletableFuture_Expecter[T]) CompleteWithError(err interface{}) *MockCompletableFuture_CompleteWithError_Call[T] {
	return &MockCompletableFuture_CompleteWithError_Call[T]{Call: _e.mock.On("CompleteWithError", err)}
}

func (_c *MockCompletableFuture_CompleteWithError_Call[T]) Run(run func(err error)) *MockCompletableFuture_CompleteWithError_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *MockCompletableFuture_CompleteWithError_Call[T]) Return() *MockCompletableFuture_CompleteWithError_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *MockCompletableFuture_CompleteWithError_Call[T]) RunAndReturn(run func(error)) *MockCompletableFuture_CompleteWithError_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields:
func (_m *MockCompletableFuture[T]) Get() T {
	ret := _m.Called()

	var r0 T
	if rf, ok := ret.Get(0).(func() T); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(T)
	}

	return r0
}

// MockCompletableFuture_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockCompletableFuture_Get_Call[T interface{}] struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *MockCompletableFuture_Expecter[T]) Get() *MockCompletableFuture_Get_Call[T] {
	return &MockCompletableFuture_Get_Call[T]{Call: _e.mock.On("Get")}
}

func (_c *MockCompletableFuture_Get_Call[T]) Run(run func()) *MockCompletableFuture_Get_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompletableFuture_Get_Call[T]) Return(_a0 T) *MockCompletableFuture_Get_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompletableFuture_Get_Call[T]) RunAndReturn(run func() T) *MockCompletableFuture_Get_Call[T] {
	_c.Call.Return(run)
	return _c
}

// GetNow provides a mock function with given fields: valueIfAbsent
func (_m *MockCompletableFuture[T]) GetNow(valueIfAbsent T) T {
	ret := _m.Called(valueIfAbsent)

	var r0 T
	if rf, ok := ret.Get(0).(func(T) T); ok {
		r0 = rf(valueIfAbsent)
	} else {
		r0 = ret.Get(0).(T)
	}

	return r0
}

// MockCompletableFuture_GetNow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNow'
type MockCompletableFuture_GetNow_Call[T interface{}] struct {
	*mock.Call
}

// GetNow is a helper method to define mock.On call
//   - valueIfAbsent T
func (_e *MockCompletableFuture_Expecter[T]) GetNow(valueIfAbsent interface{}) *MockCompletableFuture_GetNow_Call[T] {
	return &MockCompletableFuture_GetNow_Call[T]{Call: _e.mock.On("GetNow", valueIfAbsent)}
}

func (_c *MockCompletableFuture_GetNow_Call[T]) Run(run func(valueIfAbsent T)) *MockCompletableFuture_GetNow_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(T))
	})
	return _c
}

func (_c *MockCompletableFuture_GetNow_Call[T]) Return(_a0 T) *MockCompletableFuture_GetNow_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompletableFuture_GetNow_Call[T]) RunAndReturn(run func(T) T) *MockCompletableFuture_GetNow_Call[T] {
	_c.Call.Return(run)
	return _c
}

// GetWithTimeout provides a mock function with given fields: timeout
func (_m *MockCompletableFuture[T]) GetWithTimeout(timeout time.Duration) T {
	ret := _m.Called(timeout)

	var r0 T
	if rf, ok := ret.Get(0).(func(time.Duration) T); ok {
		r0 = rf(timeout)
	} else {
		r0 = ret.Get(0).(T)
	}

	return r0
}

// MockCompletableFuture_GetWithTimeout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWithTimeout'
type MockCompletableFuture_GetWithTimeout_Call[T interface{}] struct {
	*mock.Call
}

// GetWithTimeout is a helper method to define mock.On call
//   - timeout time.Duration
func (_e *MockCompletableFuture_Expecter[T]) GetWithTimeout(timeout interface{}) *MockCompletableFuture_GetWithTimeout_Call[T] {
	return &MockCompletableFuture_GetWithTimeout_Call[T]{Call: _e.mock.On("GetWithTimeout", timeout)}
}

func (_c *MockCompletableFuture_GetWithTimeout_Call[T]) Run(run func(timeout time.Duration)) *MockCompletableFuture_GetWithTimeout_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration))
	})
	return _c
}

func (_c *MockCompletableFuture_GetWithTimeout_Call[T]) Return(_a0 T) *MockCompletableFuture_GetWithTimeout_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompletableFuture_GetWithTimeout_Call[T]) RunAndReturn(run func(time.Duration) T) *MockCompletableFuture_GetWithTimeout_Call[T] {
	_c.Call.Return(run)
	return _c
}

// HandleAsync provides a mock function with given fields: _a0
func (_m *MockCompletableFuture[T]) HandleAsync(_a0 func(T, error)) {
	_m.Called(_a0)
}

// MockCompletableFuture_HandleAsync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleAsync'
type MockCompletableFuture_HandleAsync_Call[T interface{}] struct {
	*mock.Call
}

// HandleAsync is a helper method to define mock.On call
//   - _a0 func(T , error)
func (_e *MockCompletableFuture_Expecter[T]) HandleAsync(_a0 interface{}) *MockCompletableFuture_HandleAsync_Call[T] {
	return &MockCompletableFuture_HandleAsync_Call[T]{Call: _e.mock.On("HandleAsync", _a0)}
}

func (_c *MockCompletableFuture_HandleAsync_Call[T]) Run(run func(_a0 func(T, error))) *MockCompletableFuture_HandleAsync_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(T, error)))
	})
	return _c
}

func (_c *MockCompletableFuture_HandleAsync_Call[T]) Return() *MockCompletableFuture_HandleAsync_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *MockCompletableFuture_HandleAsync_Call[T]) RunAndReturn(run func(func(T, error))) *MockCompletableFuture_HandleAsync_Call[T] {
	_c.Call.Return(run)
	return _c
}

// IsCancelled provides a mock function with given fields:
func (_m *MockCompletableFuture[T]) IsCancelled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockCompletableFuture_IsCancelled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCancelled'
type MockCompletableFuture_IsCancelled_Call[T interface{}] struct {
	*mock.Call
}

// IsCancelled is a helper method to define mock.On call
func (_e *MockCompletableFuture_Expecter[T]) IsCancelled() *MockCompletableFuture_IsCancelled_Call[T] {
	return &MockCompletableFuture_IsCancelled_Call[T]{Call: _e.mock.On("IsCancelled")}
}

func (_c *MockCompletableFuture_IsCancelled_Call[T]) Run(run func()) *MockCompletableFuture_IsCancelled_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompletableFuture_IsCancelled_Call[T]) Return(_a0 bool) *MockCompletableFuture_IsCancelled_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompletableFuture_IsCancelled_Call[T]) RunAndReturn(run func() bool) *MockCompletableFuture_IsCancelled_Call[T] {
	_c.Call.Return(run)
	return _c
}

// IsCompletedWithError provides a mock function with given fields:
func (_m *MockCompletableFuture[T]) IsCompletedWithError() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockCompletableFuture_IsCompletedWithError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCompletedWithError'
type MockCompletableFuture_IsCompletedWithError_Call[T interface{}] struct {
	*mock.Call
}

// IsCompletedWithError is a helper method to define mock.On call
func (_e *MockCompletableFuture_Expecter[T]) IsCompletedWithError() *MockCompletableFuture_IsCompletedWithError_Call[T] {
	return &MockCompletableFuture_IsCompletedWithError_Call[T]{Call: _e.mock.On("IsCompletedWithError")}
}

func (_c *MockCompletableFuture_IsCompletedWithError_Call[T]) Run(run func()) *MockCompletableFuture_IsCompletedWithError_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompletableFuture_IsCompletedWithError_Call[T]) Return(_a0 bool) *MockCompletableFuture_IsCompletedWithError_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompletableFuture_IsCompletedWithError_Call[T]) RunAndReturn(run func() bool) *MockCompletableFuture_IsCompletedWithError_Call[T] {
	_c.Call.Return(run)
	return _c
}

// IsDone provides a mock function with given fields:
func (_m *MockCompletableFuture[T]) IsDone() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockCompletableFuture_IsDone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsDone'
type MockCompletableFuture_IsDone_Call[T interface{}] struct {
	*mock.Call
}

// IsDone is a helper method to define mock.On call
func (_e *MockCompletableFuture_Expecter[T]) IsDone() *MockCompletableFuture_IsDone_Call[T] {
	return &MockCompletableFuture_IsDone_Call[T]{Call: _e.mock.On("IsDone")}
}

func (_c *MockCompletableFuture_IsDone_Call[T]) Run(run func()) *MockCompletableFuture_IsDone_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCompletableFuture_IsDone_Call[T]) Return(_a0 bool) *MockCompletableFuture_IsDone_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCompletableFuture_IsDone_Call[T]) RunAndReturn(run func() bool) *MockCompletableFuture_IsDone_Call[T] {
	_c.Call.Return(run)
	return _c
}

// ThenApply provides a mock function with given fields: _a0
func (_m *MockCompletableFuture[T]) ThenApply(_a0 func(T)) {
	_m.Called(_a0)
}

// MockCompletableFuture_ThenApply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ThenApply'
type MockCompletableFuture_ThenApply_Call[T interface{}] struct {
	*mock.Call
}

// ThenApply is a helper method to define mock.On call
//   - _a0 func(T)
func (_e *MockCompletableFuture_Expecter[T]) ThenApply(_a0 interface{}) *MockCompletableFuture_ThenApply_Call[T] {
	return &MockCompletableFuture_ThenApply_Call[T]{Call: _e.mock.On("ThenApply", _a0)}
}

func (_c *MockCompletableFuture_ThenApply_Call[T]) Run(run func(_a0 func(T))) *MockCompletableFuture_ThenApply_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(T)))
	})
	return _c
}

func (_c *MockCompletableFuture_ThenApply_Call[T]) Return() *MockCompletableFuture_ThenApply_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *MockCompletableFuture_ThenApply_Call[T]) RunAndReturn(run func(func(T))) *MockCompletableFuture_ThenApply_Call[T] {
	_c.Call.Return(run)
	return _c
}

// WhenComplete provides a mock function with given fields: _a0
func (_m *MockCompletableFuture[T]) WhenComplete(_a0 func(T, error)) {
	_m.Called(_a0)
}

// MockCompletableFuture_WhenComplete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WhenComplete'
type MockCompletableFuture_WhenComplete_Call[T interface{}] struct {
	*mock.Call
}

// WhenComplete is a helper method to define mock.On call
//   - _a0 func(T , error)
func (_e *MockCompletableFuture_Expecter[T]) WhenComplete(_a0 interface{}) *MockCompletableFuture_WhenComplete_Call[T] {
	return &MockCompletableFuture_WhenComplete_Call[T]{Call: _e.mock.On("WhenComplete", _a0)}
}

func (_c *MockCompletableFuture_WhenComplete_Call[T]) Run(run func(_a0 func(T, error))) *MockCompletableFuture_WhenComplete_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(T, error)))
	})
	return _c
}

func (_c *MockCompletableFuture_WhenComplete_Call[T]) Return() *MockCompletableFuture_WhenComplete_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *MockCompletableFuture_WhenComplete_Call[T]) RunAndReturn(run func(func(T, error))) *MockCompletableFuture_WhenComplete_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewMockCompletableFuture creates a new instance of MockCompletableFuture. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCompletableFuture[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCompletableFuture[T] {
	mock := &MockCompletableFuture[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
