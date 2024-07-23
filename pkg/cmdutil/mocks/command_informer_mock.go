// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	cmdutil "github.com/skuid/skuid-cli/pkg/cmdutil"
	cobra "github.com/spf13/cobra"

	logging "github.com/skuid/skuid-cli/pkg/logging"

	mock "github.com/stretchr/testify/mock"
)

// CommandInformer is an autogenerated mock type for the CommandInformer type
type CommandInformer struct {
	mock.Mock
}

type CommandInformer_Expecter struct {
	mock *mock.Mock
}

func (_m *CommandInformer) EXPECT() *CommandInformer_Expecter {
	return &CommandInformer_Expecter{mock: &_m.Mock}
}

// AddFlags provides a mock function with given fields: cmd, flags
func (_m *CommandInformer) AddFlags(cmd *cobra.Command, flags *cmdutil.CommandFlags) {
	_m.Called(cmd, flags)
}

// CommandInformer_AddFlags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFlags'
type CommandInformer_AddFlags_Call struct {
	*mock.Call
}

// AddFlags is a helper method to define mock.On call
//   - cmd *cobra.Command
//   - flags *cmdutil.CommandFlags
func (_e *CommandInformer_Expecter) AddFlags(cmd interface{}, flags interface{}) *CommandInformer_AddFlags_Call {
	return &CommandInformer_AddFlags_Call{Call: _e.mock.On("AddFlags", cmd, flags)}
}

func (_c *CommandInformer_AddFlags_Call) Run(run func(cmd *cobra.Command, flags *cmdutil.CommandFlags)) *CommandInformer_AddFlags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*cobra.Command), args[1].(*cmdutil.CommandFlags))
	})
	return _c
}

func (_c *CommandInformer_AddFlags_Call) Return() *CommandInformer_AddFlags_Call {
	_c.Call.Return()
	return _c
}

func (_c *CommandInformer_AddFlags_Call) RunAndReturn(run func(*cobra.Command, *cmdutil.CommandFlags)) *CommandInformer_AddFlags_Call {
	_c.Call.Return(run)
	return _c
}

// ApplyEnvVars provides a mock function with given fields: cmd
func (_m *CommandInformer) ApplyEnvVars(cmd *cobra.Command) error {
	ret := _m.Called(cmd)

	if len(ret) == 0 {
		panic("no return value specified for ApplyEnvVars")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*cobra.Command) error); ok {
		r0 = rf(cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CommandInformer_ApplyEnvVars_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyEnvVars'
type CommandInformer_ApplyEnvVars_Call struct {
	*mock.Call
}

// ApplyEnvVars is a helper method to define mock.On call
//   - cmd *cobra.Command
func (_e *CommandInformer_Expecter) ApplyEnvVars(cmd interface{}) *CommandInformer_ApplyEnvVars_Call {
	return &CommandInformer_ApplyEnvVars_Call{Call: _e.mock.On("ApplyEnvVars", cmd)}
}

func (_c *CommandInformer_ApplyEnvVars_Call) Run(run func(cmd *cobra.Command)) *CommandInformer_ApplyEnvVars_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*cobra.Command))
	})
	return _c
}

func (_c *CommandInformer_ApplyEnvVars_Call) Return(_a0 error) *CommandInformer_ApplyEnvVars_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CommandInformer_ApplyEnvVars_Call) RunAndReturn(run func(*cobra.Command) error) *CommandInformer_ApplyEnvVars_Call {
	_c.Call.Return(run)
	return _c
}

// MarkFlagsMutuallyExclusive provides a mock function with given fields: cmd, flagGroups
func (_m *CommandInformer) MarkFlagsMutuallyExclusive(cmd *cobra.Command, flagGroups [][]string) {
	_m.Called(cmd, flagGroups)
}

// CommandInformer_MarkFlagsMutuallyExclusive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkFlagsMutuallyExclusive'
type CommandInformer_MarkFlagsMutuallyExclusive_Call struct {
	*mock.Call
}

// MarkFlagsMutuallyExclusive is a helper method to define mock.On call
//   - cmd *cobra.Command
//   - flagGroups [][]string
func (_e *CommandInformer_Expecter) MarkFlagsMutuallyExclusive(cmd interface{}, flagGroups interface{}) *CommandInformer_MarkFlagsMutuallyExclusive_Call {
	return &CommandInformer_MarkFlagsMutuallyExclusive_Call{Call: _e.mock.On("MarkFlagsMutuallyExclusive", cmd, flagGroups)}
}

func (_c *CommandInformer_MarkFlagsMutuallyExclusive_Call) Run(run func(cmd *cobra.Command, flagGroups [][]string)) *CommandInformer_MarkFlagsMutuallyExclusive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*cobra.Command), args[1].([][]string))
	})
	return _c
}

func (_c *CommandInformer_MarkFlagsMutuallyExclusive_Call) Return() *CommandInformer_MarkFlagsMutuallyExclusive_Call {
	_c.Call.Return()
	return _c
}

func (_c *CommandInformer_MarkFlagsMutuallyExclusive_Call) RunAndReturn(run func(*cobra.Command, [][]string)) *CommandInformer_MarkFlagsMutuallyExclusive_Call {
	_c.Call.Return(run)
	return _c
}

// SetupLogging provides a mock function with given fields: cmd, li
func (_m *CommandInformer) SetupLogging(cmd *cobra.Command, li logging.LogInformer) error {
	ret := _m.Called(cmd, li)

	if len(ret) == 0 {
		panic("no return value specified for SetupLogging")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*cobra.Command, logging.LogInformer) error); ok {
		r0 = rf(cmd, li)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CommandInformer_SetupLogging_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetupLogging'
type CommandInformer_SetupLogging_Call struct {
	*mock.Call
}

// SetupLogging is a helper method to define mock.On call
//   - cmd *cobra.Command
//   - li logging.LogInformer
func (_e *CommandInformer_Expecter) SetupLogging(cmd interface{}, li interface{}) *CommandInformer_SetupLogging_Call {
	return &CommandInformer_SetupLogging_Call{Call: _e.mock.On("SetupLogging", cmd, li)}
}

func (_c *CommandInformer_SetupLogging_Call) Run(run func(cmd *cobra.Command, li logging.LogInformer)) *CommandInformer_SetupLogging_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*cobra.Command), args[1].(logging.LogInformer))
	})
	return _c
}

func (_c *CommandInformer_SetupLogging_Call) Return(_a0 error) *CommandInformer_SetupLogging_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CommandInformer_SetupLogging_Call) RunAndReturn(run func(*cobra.Command, logging.LogInformer) error) *CommandInformer_SetupLogging_Call {
	_c.Call.Return(run)
	return _c
}

// NewCommandInformer creates a new instance of CommandInformer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommandInformer(t interface {
	mock.TestingT
	Cleanup(func())
}) *CommandInformer {
	mock := &CommandInformer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
