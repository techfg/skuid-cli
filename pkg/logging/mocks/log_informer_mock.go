// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	logging "github.com/skuid/skuid-cli/pkg/logging"
	mock "github.com/stretchr/testify/mock"
)

// LogInformer is an autogenerated mock type for the LogInformer type
type LogInformer struct {
	mock.Mock
}

type LogInformer_Expecter struct {
	mock *mock.Mock
}

func (_m *LogInformer) EXPECT() *LogInformer_Expecter {
	return &LogInformer_Expecter{mock: &_m.Mock}
}

// Setup provides a mock function with given fields: _a0
func (_m *LogInformer) Setup(_a0 logging.LoggingOptions) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Setup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(logging.LoggingOptions) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LogInformer_Setup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Setup'
type LogInformer_Setup_Call struct {
	*mock.Call
}

// Setup is a helper method to define mock.On call
//   - _a0 logging.LoggingOptions
func (_e *LogInformer_Expecter) Setup(_a0 interface{}) *LogInformer_Setup_Call {
	return &LogInformer_Setup_Call{Call: _e.mock.On("Setup", _a0)}
}

func (_c *LogInformer_Setup_Call) Run(run func(_a0 logging.LoggingOptions)) *LogInformer_Setup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(logging.LoggingOptions))
	})
	return _c
}

func (_c *LogInformer_Setup_Call) Return(_a0 error) *LogInformer_Setup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LogInformer_Setup_Call) RunAndReturn(run func(logging.LoggingOptions) error) *LogInformer_Setup_Call {
	_c.Call.Return(run)
	return _c
}

// Teardown provides a mock function with given fields:
func (_m *LogInformer) Teardown() {
	_m.Called()
}

// LogInformer_Teardown_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Teardown'
type LogInformer_Teardown_Call struct {
	*mock.Call
}

// Teardown is a helper method to define mock.On call
func (_e *LogInformer_Expecter) Teardown() *LogInformer_Teardown_Call {
	return &LogInformer_Teardown_Call{Call: _e.mock.On("Teardown")}
}

func (_c *LogInformer_Teardown_Call) Run(run func()) *LogInformer_Teardown_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *LogInformer_Teardown_Call) Return() *LogInformer_Teardown_Call {
	_c.Call.Return()
	return _c
}

func (_c *LogInformer_Teardown_Call) RunAndReturn(run func()) *LogInformer_Teardown_Call {
	_c.Call.Return(run)
	return _c
}

// NewLogInformer creates a new instance of LogInformer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogInformer(t interface {
	mock.TestingT
	Cleanup(func())
}) *LogInformer {
	mock := &LogInformer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
