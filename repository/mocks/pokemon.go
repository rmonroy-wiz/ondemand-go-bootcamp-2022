// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	model "github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// PokemonRepositoryMock is an autogenerated mock type for the PokemonRepository type
type PokemonRepositoryMock struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *PokemonRepositoryMock) GetAll() ([]*model.PokemonDTO, *model.ErrorHandler) {
	ret := _m.Called()

	var r0 []*model.PokemonDTO
	if rf, ok := ret.Get(0).(func() []*model.PokemonDTO); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PokemonDTO)
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
func (_m *PokemonRepositoryMock) GetByID(id int) (*model.PokemonDTO, *model.ErrorHandler) {
	ret := _m.Called(id)

	var r0 *model.PokemonDTO
	if rf, ok := ret.Get(0).(func(int) *model.PokemonDTO); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PokemonDTO)
		}
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

// GetCSVDataInMemory provides a mock function with given fields:
func (_m *PokemonRepositoryMock) GetCSVDataInMemory() (map[int]model.PokemonCSV, *model.ErrorHandler) {
	ret := _m.Called()

	var r0 map[int]model.PokemonCSV
	if rf, ok := ret.Get(0).(func() map[int]model.PokemonCSV); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int]model.PokemonCSV)
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

// StoreToCSV provides a mock function with given fields: pokemonAPI
func (_m *PokemonRepositoryMock) StoreToCSV(pokemonAPI model.PokemonAPI) (*model.PokemonDTO, *model.ErrorHandler) {
	ret := _m.Called(pokemonAPI)

	var r0 *model.PokemonDTO
	if rf, ok := ret.Get(0).(func(model.PokemonAPI) *model.PokemonDTO); ok {
		r0 = rf(pokemonAPI)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PokemonDTO)
		}
	}

	var r1 *model.ErrorHandler
	if rf, ok := ret.Get(1).(func(model.PokemonAPI) *model.ErrorHandler); ok {
		r1 = rf(pokemonAPI)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.ErrorHandler)
		}
	}

	return r0, r1
}

// NewPokemonRepositoryMock creates a new instance of PokemonRepositoryMock. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewPokemonRepositoryMock(t testing.TB) *PokemonRepositoryMock {
	mock := &PokemonRepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}