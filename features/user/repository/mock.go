package repository

import (
	"completeUser/features/user/domain"

	"github.com/stretchr/testify/mock"
)

type RepoQueryMock struct {
	mock.Mock
}

func (rq *RepoQueryMock) Insert(newUser domain.Core) (domain.Core, error) {
	ret := rq.Called(newUser)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (rq *RepoQueryMock) Update(updatedData domain.Core) (domain.Core, error) {
	ret := rq.Called(updatedData)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(updatedData)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(updatedData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (rq *RepoQueryMock) Delete(ID uint) error {
	ret := rq.Called(ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(int(ID))
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (rq *RepoQueryMock) Get(ID uint) (domain.Core, error) {
	ret := rq.Called(ID)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(int) domain.Core); ok {
		r0 = rf(int(ID))
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(int(ID))
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (rq *RepoQueryMock) GetAll() ([]domain.Core, error) {
	ret := rq.Called()

	var r0 []domain.Core
	if rf, ok := ret.Get(0).(func() []domain.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
