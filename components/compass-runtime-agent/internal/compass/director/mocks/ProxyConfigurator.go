// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	cache "github.com/kyma-project/kyma/components/compass-runtime-agent/internal/compass/cache"

	mock "github.com/stretchr/testify/mock"
)

// ProxyConfigurator is an autogenerated mock type for the ProxyConfigurator type
type ProxyConfigurator struct {
	mock.Mock
}

// SetURLAndCerts provides a mock function with given fields: data
func (_m *ProxyConfigurator) SetURLAndCerts(data cache.ConnectionData) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(cache.ConnectionData) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
