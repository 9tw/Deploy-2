package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID      uint
	Email   string
	Name    string
	Phone   string
	Address string
}

type Repository interface {
	Insert(newUser Core) (Core, error)
	Update(updatedData Core) (Core, error)
	Delete(ID uint) error
	Get(ID uint) (Core, error)
	GetAll() ([]Core, error)
}

type Service interface {
	AddUser(newUser Core) (Core, error)
	UpdateB(updatedData Core) (Core, error)
	Profile(ID uint) (Core, error)
	ShowAllUser() ([]Core, error)
	Delete(ID uint) error
}

type Handler interface {
	AddUser() echo.HandlerFunc
	ShowAllUser() echo.HandlerFunc
	Delete(ID uint) echo.HandlerFunc
	UpdateBook() echo.HandlerFunc
	OneUser(ID uint) echo.HandlerFunc
}
