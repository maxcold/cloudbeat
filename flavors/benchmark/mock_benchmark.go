// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by mockery v2.24.0. DO NOT EDIT.

package benchmark

import (
	context "context"

	config "github.com/elastic/cloudbeat/config"

	dataprovider "github.com/elastic/cloudbeat/dataprovider"

	fetching "github.com/elastic/cloudbeat/resources/fetching"

	logp "github.com/elastic/elastic-agent-libs/logp"

	mock "github.com/stretchr/testify/mock"

	registry "github.com/elastic/cloudbeat/resources/fetching/registry"
)

// MockBenchmark is an autogenerated mock type for the Benchmark type
type MockBenchmark struct {
	mock.Mock
}

type MockBenchmark_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBenchmark) EXPECT() *MockBenchmark_Expecter {
	return &MockBenchmark_Expecter{mock: &_m.Mock}
}

// Initialize provides a mock function with given fields: ctx, log, cfg, ch
func (_m *MockBenchmark) Initialize(ctx context.Context, log *logp.Logger, cfg *config.Config, ch chan fetching.ResourceInfo) (registry.Registry, dataprovider.CommonDataProvider, dataprovider.IdProvider, error) {
	ret := _m.Called(ctx, log, cfg, ch)

	var r0 registry.Registry
	var r1 dataprovider.CommonDataProvider
	var r2 dataprovider.IdProvider
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, *logp.Logger, *config.Config, chan fetching.ResourceInfo) (registry.Registry, dataprovider.CommonDataProvider, dataprovider.IdProvider, error)); ok {
		return rf(ctx, log, cfg, ch)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *logp.Logger, *config.Config, chan fetching.ResourceInfo) registry.Registry); ok {
		r0 = rf(ctx, log, cfg, ch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(registry.Registry)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *logp.Logger, *config.Config, chan fetching.ResourceInfo) dataprovider.CommonDataProvider); ok {
		r1 = rf(ctx, log, cfg, ch)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(dataprovider.CommonDataProvider)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *logp.Logger, *config.Config, chan fetching.ResourceInfo) dataprovider.IdProvider); ok {
		r2 = rf(ctx, log, cfg, ch)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(dataprovider.IdProvider)
		}
	}

	if rf, ok := ret.Get(3).(func(context.Context, *logp.Logger, *config.Config, chan fetching.ResourceInfo) error); ok {
		r3 = rf(ctx, log, cfg, ch)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// MockBenchmark_Initialize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Initialize'
type MockBenchmark_Initialize_Call struct {
	*mock.Call
}

// Initialize is a helper method to define mock.On call
//   - ctx context.Context
//   - log *logp.Logger
//   - cfg *config.Config
//   - ch chan fetching.ResourceInfo
func (_e *MockBenchmark_Expecter) Initialize(ctx interface{}, log interface{}, cfg interface{}, ch interface{}) *MockBenchmark_Initialize_Call {
	return &MockBenchmark_Initialize_Call{Call: _e.mock.On("Initialize", ctx, log, cfg, ch)}
}

func (_c *MockBenchmark_Initialize_Call) Run(run func(ctx context.Context, log *logp.Logger, cfg *config.Config, ch chan fetching.ResourceInfo)) *MockBenchmark_Initialize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*logp.Logger), args[2].(*config.Config), args[3].(chan fetching.ResourceInfo))
	})
	return _c
}

func (_c *MockBenchmark_Initialize_Call) Return(_a0 registry.Registry, _a1 dataprovider.CommonDataProvider, _a2 dataprovider.IdProvider, _a3 error) *MockBenchmark_Initialize_Call {
	_c.Call.Return(_a0, _a1, _a2, _a3)
	return _c
}

func (_c *MockBenchmark_Initialize_Call) RunAndReturn(run func(context.Context, *logp.Logger, *config.Config, chan fetching.ResourceInfo) (registry.Registry, dataprovider.CommonDataProvider, dataprovider.IdProvider, error)) *MockBenchmark_Initialize_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: ctx
func (_m *MockBenchmark) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBenchmark_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockBenchmark_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockBenchmark_Expecter) Run(ctx interface{}) *MockBenchmark_Run_Call {
	return &MockBenchmark_Run_Call{Call: _e.mock.On("Run", ctx)}
}

func (_c *MockBenchmark_Run_Call) Run(run func(ctx context.Context)) *MockBenchmark_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockBenchmark_Run_Call) Return(_a0 error) *MockBenchmark_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBenchmark_Run_Call) RunAndReturn(run func(context.Context) error) *MockBenchmark_Run_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockBenchmark) Stop() {
	_m.Called()
}

// MockBenchmark_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockBenchmark_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockBenchmark_Expecter) Stop() *MockBenchmark_Stop_Call {
	return &MockBenchmark_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockBenchmark_Stop_Call) Run(run func()) *MockBenchmark_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBenchmark_Stop_Call) Return() *MockBenchmark_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBenchmark_Stop_Call) RunAndReturn(run func()) *MockBenchmark_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// checkDependencies provides a mock function with given fields:
func (_m *MockBenchmark) checkDependencies() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBenchmark_checkDependencies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'checkDependencies'
type MockBenchmark_checkDependencies_Call struct {
	*mock.Call
}

// checkDependencies is a helper method to define mock.On call
func (_e *MockBenchmark_Expecter) checkDependencies() *MockBenchmark_checkDependencies_Call {
	return &MockBenchmark_checkDependencies_Call{Call: _e.mock.On("checkDependencies")}
}

func (_c *MockBenchmark_checkDependencies_Call) Run(run func()) *MockBenchmark_checkDependencies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBenchmark_checkDependencies_Call) Return(_a0 error) *MockBenchmark_checkDependencies_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBenchmark_checkDependencies_Call) RunAndReturn(run func() error) *MockBenchmark_checkDependencies_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockBenchmark interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockBenchmark creates a new instance of MockBenchmark. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockBenchmark(t mockConstructorTestingTNewMockBenchmark) *MockBenchmark {
	mock := &MockBenchmark{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
