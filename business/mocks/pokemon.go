// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	testing "testing"

	model "github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
	mock "github.com/stretchr/testify/mock"
)

// PokemonBusinessMock is an autogenerated mock type for the PokemonBusiness type
type PokemonBusinessMock struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *PokemonBusinessMock) GetAll() ([]model.PokemonDTO, *model.ErrorHandler) {
	ret := _m.Called()

	var r0 []model.PokemonDTO
	if rf, ok := ret.Get(0).(func() []model.PokemonDTO); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PokemonDTO)
		}
	}

	var r1 *model.ErrorHandler
	if rf, ok := ret.Get(1).(func() *model.ErrorHandler); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.ErrorHandler)
		}
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *PokemonBusinessMock) GetByID(id int) (model.PokemonDTO, *model.ErrorHandler) {
	ret := _m.Called(id)

	var r0 model.PokemonDTO
	if rf, ok := ret.Get(0).(func(int) model.PokemonDTO); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.PokemonDTO)
	}

	var r1 *model.ErrorHandler
	if rf, ok := ret.Get(1).(func(int) *model.ErrorHandler); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.ErrorHandler)
		}
	}

	return r0, r1
}

// SearchPokemon provides a mock function with given fields: typeSearch, items, itemsPerWorker
func (_m *PokemonBusinessMock) SearchPokemon(typeSearch string, items int, itemsPerWorker int) ([]model.PokemonDTO, *model.ErrorHandler) {
	ret := _m.Called(typeSearch, items, itemsPerWorker)

	var r0 []model.PokemonDTO
	if rf, ok := ret.Get(0).(func(string, int, int) []model.PokemonDTO); ok {
		r0 = rf(typeSearch, items, itemsPerWorker)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PokemonDTO)
		}
	}

	var r1 *model.ErrorHandler
	if rf, ok := ret.Get(1).(func(string, int, int) *model.ErrorHandler); ok {
		r1 = rf(typeSearch, items, itemsPerWorker)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.ErrorHandler)
		}
	}

	return r0, r1
}

// StoreByID provides a mock function with given fields: id
func (_m *PokemonBusinessMock) StoreByID(id int) (model.PokemonDTO, *model.ErrorHandler) {
	ret := _m.Called(id)

	var r0 model.PokemonDTO
	if rf, ok := ret.Get(0).(func(int) model.PokemonDTO); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.PokemonDTO)
	}

	var r1 *model.ErrorHandler
	if rf, ok := ret.Get(1).(func(int) *model.ErrorHandler); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.ErrorHandler)
		}
	}

	return r0, r1
}

// NewPokemonBusinessMock creates a new instance of PokemonBusinessMock. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewPokemonBusinessMock(t testing.TB) *PokemonBusinessMock {
	mock := &PokemonBusinessMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
