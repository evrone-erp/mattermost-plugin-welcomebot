// Code generated by mockery v2.53.0. DO NOT EDIT.

package usecase

import (
	internalmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	mock "github.com/stretchr/testify/mock"

	model "github.com/mattermost/mattermost/server/public/model"
)

// MockChannelWelcomeRepo is an autogenerated mock type for the ChannelWelcomeRepo type
type MockChannelWelcomeRepo struct {
	mock.Mock
}

type MockChannelWelcomeRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockChannelWelcomeRepo) EXPECT() *MockChannelWelcomeRepo_Expecter {
	return &MockChannelWelcomeRepo_Expecter{mock: &_m.Mock}
}

// DeletePersonalChanelWelcome provides a mock function with given fields: channelID
func (_m *MockChannelWelcomeRepo) DeletePersonalChanelWelcome(channelID string) *model.AppError {
	ret := _m.Called(channelID)

	if len(ret) == 0 {
		panic("no return value specified for DeletePersonalChanelWelcome")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// MockChannelWelcomeRepo_DeletePersonalChanelWelcome_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeletePersonalChanelWelcome'
type MockChannelWelcomeRepo_DeletePersonalChanelWelcome_Call struct {
	*mock.Call
}

// DeletePersonalChanelWelcome is a helper method to define mock.On call
//   - channelID string
func (_e *MockChannelWelcomeRepo_Expecter) DeletePersonalChanelWelcome(channelID interface{}) *MockChannelWelcomeRepo_DeletePersonalChanelWelcome_Call {
	return &MockChannelWelcomeRepo_DeletePersonalChanelWelcome_Call{Call: _e.mock.On("DeletePersonalChanelWelcome", channelID)}
}

func (_c *MockChannelWelcomeRepo_DeletePersonalChanelWelcome_Call) Run(run func(channelID string)) *MockChannelWelcomeRepo_DeletePersonalChanelWelcome_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockChannelWelcomeRepo_DeletePersonalChanelWelcome_Call) Return(_a0 *model.AppError) *MockChannelWelcomeRepo_DeletePersonalChanelWelcome_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChannelWelcomeRepo_DeletePersonalChanelWelcome_Call) RunAndReturn(run func(string) *model.AppError) *MockChannelWelcomeRepo_DeletePersonalChanelWelcome_Call {
	_c.Call.Return(run)
	return _c
}

// DeletePublishedChanelWelcome provides a mock function with given fields: channelID
func (_m *MockChannelWelcomeRepo) DeletePublishedChanelWelcome(channelID string) *model.AppError {
	ret := _m.Called(channelID)

	if len(ret) == 0 {
		panic("no return value specified for DeletePublishedChanelWelcome")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// MockChannelWelcomeRepo_DeletePublishedChanelWelcome_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeletePublishedChanelWelcome'
type MockChannelWelcomeRepo_DeletePublishedChanelWelcome_Call struct {
	*mock.Call
}

// DeletePublishedChanelWelcome is a helper method to define mock.On call
//   - channelID string
func (_e *MockChannelWelcomeRepo_Expecter) DeletePublishedChanelWelcome(channelID interface{}) *MockChannelWelcomeRepo_DeletePublishedChanelWelcome_Call {
	return &MockChannelWelcomeRepo_DeletePublishedChanelWelcome_Call{Call: _e.mock.On("DeletePublishedChanelWelcome", channelID)}
}

func (_c *MockChannelWelcomeRepo_DeletePublishedChanelWelcome_Call) Run(run func(channelID string)) *MockChannelWelcomeRepo_DeletePublishedChanelWelcome_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockChannelWelcomeRepo_DeletePublishedChanelWelcome_Call) Return(_a0 *model.AppError) *MockChannelWelcomeRepo_DeletePublishedChanelWelcome_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChannelWelcomeRepo_DeletePublishedChanelWelcome_Call) RunAndReturn(run func(string) *model.AppError) *MockChannelWelcomeRepo_DeletePublishedChanelWelcome_Call {
	_c.Call.Return(run)
	return _c
}

// GetPersonalChanelWelcome provides a mock function with given fields: channelID
func (_m *MockChannelWelcomeRepo) GetPersonalChanelWelcome(channelID string) (*internalmodel.ChannelWelcome, *model.AppError) {
	ret := _m.Called(channelID)

	if len(ret) == 0 {
		panic("no return value specified for GetPersonalChanelWelcome")
	}

	var r0 *internalmodel.ChannelWelcome
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string) (*internalmodel.ChannelWelcome, *model.AppError)); ok {
		return rf(channelID)
	}
	if rf, ok := ret.Get(0).(func(string) *internalmodel.ChannelWelcome); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalmodel.ChannelWelcome)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(channelID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// MockChannelWelcomeRepo_GetPersonalChanelWelcome_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPersonalChanelWelcome'
type MockChannelWelcomeRepo_GetPersonalChanelWelcome_Call struct {
	*mock.Call
}

// GetPersonalChanelWelcome is a helper method to define mock.On call
//   - channelID string
func (_e *MockChannelWelcomeRepo_Expecter) GetPersonalChanelWelcome(channelID interface{}) *MockChannelWelcomeRepo_GetPersonalChanelWelcome_Call {
	return &MockChannelWelcomeRepo_GetPersonalChanelWelcome_Call{Call: _e.mock.On("GetPersonalChanelWelcome", channelID)}
}

func (_c *MockChannelWelcomeRepo_GetPersonalChanelWelcome_Call) Run(run func(channelID string)) *MockChannelWelcomeRepo_GetPersonalChanelWelcome_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockChannelWelcomeRepo_GetPersonalChanelWelcome_Call) Return(_a0 *internalmodel.ChannelWelcome, _a1 *model.AppError) *MockChannelWelcomeRepo_GetPersonalChanelWelcome_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChannelWelcomeRepo_GetPersonalChanelWelcome_Call) RunAndReturn(run func(string) (*internalmodel.ChannelWelcome, *model.AppError)) *MockChannelWelcomeRepo_GetPersonalChanelWelcome_Call {
	_c.Call.Return(run)
	return _c
}

// GetPublishedChanelWelcome provides a mock function with given fields: channelID
func (_m *MockChannelWelcomeRepo) GetPublishedChanelWelcome(channelID string) (*internalmodel.ChannelWelcome, *model.AppError) {
	ret := _m.Called(channelID)

	if len(ret) == 0 {
		panic("no return value specified for GetPublishedChanelWelcome")
	}

	var r0 *internalmodel.ChannelWelcome
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string) (*internalmodel.ChannelWelcome, *model.AppError)); ok {
		return rf(channelID)
	}
	if rf, ok := ret.Get(0).(func(string) *internalmodel.ChannelWelcome); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internalmodel.ChannelWelcome)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(channelID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// MockChannelWelcomeRepo_GetPublishedChanelWelcome_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPublishedChanelWelcome'
type MockChannelWelcomeRepo_GetPublishedChanelWelcome_Call struct {
	*mock.Call
}

// GetPublishedChanelWelcome is a helper method to define mock.On call
//   - channelID string
func (_e *MockChannelWelcomeRepo_Expecter) GetPublishedChanelWelcome(channelID interface{}) *MockChannelWelcomeRepo_GetPublishedChanelWelcome_Call {
	return &MockChannelWelcomeRepo_GetPublishedChanelWelcome_Call{Call: _e.mock.On("GetPublishedChanelWelcome", channelID)}
}

func (_c *MockChannelWelcomeRepo_GetPublishedChanelWelcome_Call) Run(run func(channelID string)) *MockChannelWelcomeRepo_GetPublishedChanelWelcome_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockChannelWelcomeRepo_GetPublishedChanelWelcome_Call) Return(_a0 *internalmodel.ChannelWelcome, _a1 *model.AppError) *MockChannelWelcomeRepo_GetPublishedChanelWelcome_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChannelWelcomeRepo_GetPublishedChanelWelcome_Call) RunAndReturn(run func(string) (*internalmodel.ChannelWelcome, *model.AppError)) *MockChannelWelcomeRepo_GetPublishedChanelWelcome_Call {
	_c.Call.Return(run)
	return _c
}

// ListChannelsWithWelcome provides a mock function with no fields
func (_m *MockChannelWelcomeRepo) ListChannelsWithWelcome() ([]string, []string, *model.AppError) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListChannelsWithWelcome")
	}

	var r0 []string
	var r1 []string
	var r2 *model.AppError
	if rf, ok := ret.Get(0).(func() ([]string, []string, *model.AppError)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func() []string); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	if rf, ok := ret.Get(2).(func() *model.AppError); ok {
		r2 = rf()
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*model.AppError)
		}
	}

	return r0, r1, r2
}

// MockChannelWelcomeRepo_ListChannelsWithWelcome_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListChannelsWithWelcome'
type MockChannelWelcomeRepo_ListChannelsWithWelcome_Call struct {
	*mock.Call
}

// ListChannelsWithWelcome is a helper method to define mock.On call
func (_e *MockChannelWelcomeRepo_Expecter) ListChannelsWithWelcome() *MockChannelWelcomeRepo_ListChannelsWithWelcome_Call {
	return &MockChannelWelcomeRepo_ListChannelsWithWelcome_Call{Call: _e.mock.On("ListChannelsWithWelcome")}
}

func (_c *MockChannelWelcomeRepo_ListChannelsWithWelcome_Call) Run(run func()) *MockChannelWelcomeRepo_ListChannelsWithWelcome_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockChannelWelcomeRepo_ListChannelsWithWelcome_Call) Return(_a0 []string, _a1 []string, _a2 *model.AppError) *MockChannelWelcomeRepo_ListChannelsWithWelcome_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockChannelWelcomeRepo_ListChannelsWithWelcome_Call) RunAndReturn(run func() ([]string, []string, *model.AppError)) *MockChannelWelcomeRepo_ListChannelsWithWelcome_Call {
	_c.Call.Return(run)
	return _c
}

// SetPersonalChanelWelcome provides a mock function with given fields: channelID, message
func (_m *MockChannelWelcomeRepo) SetPersonalChanelWelcome(channelID string, message string) *model.AppError {
	ret := _m.Called(channelID, message)

	if len(ret) == 0 {
		panic("no return value specified for SetPersonalChanelWelcome")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, string) *model.AppError); ok {
		r0 = rf(channelID, message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// MockChannelWelcomeRepo_SetPersonalChanelWelcome_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPersonalChanelWelcome'
type MockChannelWelcomeRepo_SetPersonalChanelWelcome_Call struct {
	*mock.Call
}

// SetPersonalChanelWelcome is a helper method to define mock.On call
//   - channelID string
//   - message string
func (_e *MockChannelWelcomeRepo_Expecter) SetPersonalChanelWelcome(channelID interface{}, message interface{}) *MockChannelWelcomeRepo_SetPersonalChanelWelcome_Call {
	return &MockChannelWelcomeRepo_SetPersonalChanelWelcome_Call{Call: _e.mock.On("SetPersonalChanelWelcome", channelID, message)}
}

func (_c *MockChannelWelcomeRepo_SetPersonalChanelWelcome_Call) Run(run func(channelID string, message string)) *MockChannelWelcomeRepo_SetPersonalChanelWelcome_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockChannelWelcomeRepo_SetPersonalChanelWelcome_Call) Return(_a0 *model.AppError) *MockChannelWelcomeRepo_SetPersonalChanelWelcome_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChannelWelcomeRepo_SetPersonalChanelWelcome_Call) RunAndReturn(run func(string, string) *model.AppError) *MockChannelWelcomeRepo_SetPersonalChanelWelcome_Call {
	_c.Call.Return(run)
	return _c
}

// SetPublishedChanelWelcome provides a mock function with given fields: channelID, message
func (_m *MockChannelWelcomeRepo) SetPublishedChanelWelcome(channelID string, message string) *model.AppError {
	ret := _m.Called(channelID, message)

	if len(ret) == 0 {
		panic("no return value specified for SetPublishedChanelWelcome")
	}

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, string) *model.AppError); ok {
		r0 = rf(channelID, message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// MockChannelWelcomeRepo_SetPublishedChanelWelcome_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPublishedChanelWelcome'
type MockChannelWelcomeRepo_SetPublishedChanelWelcome_Call struct {
	*mock.Call
}

// SetPublishedChanelWelcome is a helper method to define mock.On call
//   - channelID string
//   - message string
func (_e *MockChannelWelcomeRepo_Expecter) SetPublishedChanelWelcome(channelID interface{}, message interface{}) *MockChannelWelcomeRepo_SetPublishedChanelWelcome_Call {
	return &MockChannelWelcomeRepo_SetPublishedChanelWelcome_Call{Call: _e.mock.On("SetPublishedChanelWelcome", channelID, message)}
}

func (_c *MockChannelWelcomeRepo_SetPublishedChanelWelcome_Call) Run(run func(channelID string, message string)) *MockChannelWelcomeRepo_SetPublishedChanelWelcome_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockChannelWelcomeRepo_SetPublishedChanelWelcome_Call) Return(_a0 *model.AppError) *MockChannelWelcomeRepo_SetPublishedChanelWelcome_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChannelWelcomeRepo_SetPublishedChanelWelcome_Call) RunAndReturn(run func(string, string) *model.AppError) *MockChannelWelcomeRepo_SetPublishedChanelWelcome_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockChannelWelcomeRepo creates a new instance of MockChannelWelcomeRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockChannelWelcomeRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockChannelWelcomeRepo {
	mock := &MockChannelWelcomeRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
