package delivery

import (
	"completeUser/features/user/domain"
)

type RegisterFormat struct {
	Email   string `json:"email" form:"email"`
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

type UpdateFormat struct {
	ID      uint   `json:"id" form:"id"`
	Email   string `json:"email" form:"email"`
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.Core{Email: cnv.Email, Name: cnv.Name, Phone: cnv.Phone, Address: cnv.Address}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return domain.Core{ID: cnv.ID, Email: cnv.Email, Name: cnv.Name, Phone: cnv.Phone, Address: cnv.Address}
	}
	return domain.Core{}
}
