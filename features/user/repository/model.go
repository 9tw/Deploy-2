package repository

import (
	"completeUser/features/user/domain"
)

type User struct {
	ID      uint
	Email   string
	Name    string
	Phone   string
	Address string
}

func FromDomain(du domain.Core) User {
	return User{
		ID:      du.ID,
		Email:   du.Email,
		Name:    du.Name,
		Phone:   du.Phone,
		Address: du.Address,
	}
}

func ToDomain(u User) domain.Core {
	return domain.Core{
		ID:      u.ID,
		Email:   u.Email,
		Name:    u.Name,
		Phone:   u.Phone,
		Address: u.Address,
	}
}

func ToDomainArray(au []User) []domain.Core {
	var res []domain.Core
	for _, val := range au {
		res = append(res, domain.Core{ID: val.ID, Email: val.Email, Name: val.Name, Phone: val.Phone, Address: val.Address})
	}

	return res
}
