// Code generated by mockery v1.1.1. DO NOT EDIT.

package automock

import (
	git "github.com/kyma-project/kyma/components/function-controller/internal/git"
	mock "github.com/stretchr/testify/mock"
)

// GitOperator is an autogenerated mock type for the GitOperator type
type GitOperator struct {
	mock.Mock
}

// Clone provides a mock function with given fields: path, options
func (_m *GitOperator) Clone(path string, options git.Options) (string, error) {
	ret := _m.Called(path, options)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, git.Options) string); ok {
		r0 = rf(path, options)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, git.Options) error); ok {
		r1 = rf(path, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LastCommit provides a mock function with given fields: options
func (_m *GitOperator) LastCommit(options git.Options) (string, error) {
	ret := _m.Called(options)

	var r0 string
	if rf, ok := ret.Get(0).(func(git.Options) string); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(git.Options) error); ok {
		r1 = rf(options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
