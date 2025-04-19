// Code generated by mockery v2.46.3. DO NOT EDIT.

package mockproduct

import (
	entity "github.com/hanhnham91/order-service/entity"
	specifications "github.com/hanhnham91/order-service/repository/specifications"
	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// Find provides a mock function with given fields: spec
func (_m *MockRepository) Find(spec specifications.I) ([]entity.Product, error) {
	ret := _m.Called(spec)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 []entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(specifications.I) ([]entity.Product, error)); ok {
		return rf(spec)
	}
	if rf, ok := ret.Get(0).(func(specifications.I) []entity.Product); ok {
		r0 = rf(spec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(specifications.I) error); ok {
		r1 = rf(spec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_Find_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Find'
type MockRepository_Find_Call struct {
	*mock.Call
}

// Find is a helper method to define mock.On call
//   - spec specifications.I
func (_e *MockRepository_Expecter) Find(spec interface{}) *MockRepository_Find_Call {
	return &MockRepository_Find_Call{Call: _e.mock.On("Find", spec)}
}

func (_c *MockRepository_Find_Call) Run(run func(spec specifications.I)) *MockRepository_Find_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(specifications.I))
	})
	return _c
}

func (_c *MockRepository_Find_Call) Return(_a0 []entity.Product, _a1 error) *MockRepository_Find_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_Find_Call) RunAndReturn(run func(specifications.I) ([]entity.Product, error)) *MockRepository_Find_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields:
func (_m *MockRepository) FindAll() ([]entity.Product, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.Product, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Product)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type MockRepository_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
func (_e *MockRepository_Expecter) FindAll() *MockRepository_FindAll_Call {
	return &MockRepository_FindAll_Call{Call: _e.mock.On("FindAll")}
}

func (_c *MockRepository_FindAll_Call) Run(run func()) *MockRepository_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRepository_FindAll_Call) Return(_a0 []entity.Product, _a1 error) *MockRepository_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindAll_Call) RunAndReturn(run func() ([]entity.Product, error)) *MockRepository_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: spec
func (_m *MockRepository) Get(spec specifications.I) (entity.Product, error) {
	ret := _m.Called(spec)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(specifications.I) (entity.Product, error)); ok {
		return rf(spec)
	}
	if rf, ok := ret.Get(0).(func(specifications.I) entity.Product); ok {
		r0 = rf(spec)
	} else {
		r0 = ret.Get(0).(entity.Product)
	}

	if rf, ok := ret.Get(1).(func(specifications.I) error); ok {
		r1 = rf(spec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockRepository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - spec specifications.I
func (_e *MockRepository_Expecter) Get(spec interface{}) *MockRepository_Get_Call {
	return &MockRepository_Get_Call{Call: _e.mock.On("Get", spec)}
}

func (_c *MockRepository_Get_Call) Run(run func(spec specifications.I)) *MockRepository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(specifications.I))
	})
	return _c
}

func (_c *MockRepository_Get_Call) Return(_a0 entity.Product, _a1 error) *MockRepository_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_Get_Call) RunAndReturn(run func(specifications.I) (entity.Product, error)) *MockRepository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
