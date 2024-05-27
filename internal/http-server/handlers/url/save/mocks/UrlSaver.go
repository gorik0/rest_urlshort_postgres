// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UrlSaver is an autogenerated mock type for the UrlSaver type
type UrlSaver struct {
	mock.Mock
}

// SaveUrl provides a mock function with given fields: urlTYoSave, alias
func (_m *UrlSaver) SaveUrl(urlTYoSave string, alias string) (int64, error) {
	ret := _m.Called(urlTYoSave, alias)

	if len(ret) == 0 {
		panic("no return value specified for SaveUrl")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (int64, error)); ok {
		return rf(urlTYoSave, alias)
	}
	if rf, ok := ret.Get(0).(func(string, string) int64); ok {
		r0 = rf(urlTYoSave, alias)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(urlTYoSave, alias)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUrlSaver creates a new instance of UrlSaver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUrlSaver(t interface {
	mock.TestingT
	Cleanup(func())
}) *UrlSaver {
	mock := &UrlSaver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
