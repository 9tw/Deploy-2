package services

import (
	dom "completeUser/features/user/domain"
	rep "completeUser/features/user/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var queryRepo = &rep.RepoQueryMock{Mock: mock.Mock{}}
var serviceRepo = UserService{qry: queryRepo}

func TestAddUser(t *testing.T) {
	repo := new(rep.RepoQueryMock)
	insertData := dom.Core{Email: "geger@gmail.com", Name: "Gerry", Phone: "1234", Address: "Malang"}
	returnData := dom.Core{ID: uint(1), Email: "geger@gmail.com", Name: "Gerry", Phone: "1234", Address: "Malang"}

	t.Run("Success", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)

		res, err := srv.AddUser(insertData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, returnData.Email, res.Email)
		assert.Equal(t, returnData.Name, res.Name)
		assert.Equal(t, returnData.Phone, res.Phone)
		assert.Equal(t, returnData.Address, res.Address)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(dom.Core{}, errors.New("some problem on database")).Once()
		srv := New(repo)

		res, err := srv.AddUser(dom.Core{})
		assert.Equal(t, uint(0), res.ID)
		assert.Equal(t, "", res.Name)
		assert.Error(t, err)
		assert.EqualError(t, err, "some problem on database")
		repo.AssertExpectations(t)
	})
}

func TestShowAllUser(t *testing.T) {
	repo := new(rep.RepoQueryMock)
	returnData := []dom.Core{{ID: uint(1), Email: "geger@gmail.com", Name: "Gerry", Phone: "1234", Address: "Malang"}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll").Return(returnData, nil).Once()

		srv := New(repo)

		res, err := srv.ShowAllUser()
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get All", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("data not found")).Once()

		srv := New(repo)

		res, err := srv.ShowAllUser()
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestProfile(t *testing.T) {
	repo := new(rep.RepoQueryMock)
	returnData := dom.Core{ID: uint(1), Email: "geger@gmail.com", Name: "Gerry", Phone: "1234", Address: "Malang"}

	t.Run("Success", func(t *testing.T) {
		repo.On("Get", uint(1)).Return(returnData, nil).Once()

		srv := New(repo)

		_, err := srv.Profile(uint(1))
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Out of Index", func(t *testing.T) {
		repo.On("Get", uint(2)).Return(dom.Core{}, errors.New("error out of index")).Once()

		srv := New(repo)

		_, err := srv.Profile(uint(2))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdateB(t *testing.T) {
	repo := new(rep.RepoQueryMock)
	updateData := dom.Core{Email: "geger@gmail.com", Name: "Gerry", Phone: "1234", Address: "Malang"}
	returnData := dom.Core{ID: uint(1), Email: "geger@gmail.com", Name: "Gerry", Phone: "1234", Address: "Malang"}

	t.Run("Success", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)

		res, err := srv.UpdateB(updateData)
		assert.NoError(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, returnData.Email, res.Email)
		assert.Equal(t, returnData.Name, res.Name)
		assert.Equal(t, returnData.Phone, res.Phone)
		assert.Equal(t, returnData.Address, res.Address)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(dom.Core{}, errors.New("some problem on database")).Once()
		srv := New(repo)

		res, err := srv.UpdateB(dom.Core{})
		assert.Equal(t, uint(0), res.ID)
		assert.Equal(t, "", res.Name)
		assert.Error(t, err)
		assert.EqualError(t, err, "some problem on database")
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(rep.RepoQueryMock)
	//returnData := dom.Core{ID: uint(1), Email: "geger@gmail.com", Name: "Gerry", Phone: "1234", Address: "Malang"}

	t.Run("Success", func(t *testing.T) {
		repo.On("Delete", uint(1)).Return(nil).Once()

		srv := New(repo)

		err := srv.Delete(uint(1))
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Out of Index", func(t *testing.T) {
		repo.On("Delete", uint(2)).Return(errors.New("error out of index")).Once()

		srv := New(repo)

		err := srv.Delete(uint(2))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}
