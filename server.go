package main

import (
	"completeUser/config"
	ud "completeUser/features/user/delivery"
	ur "completeUser/features/user/repository"
	us "completeUser/features/user/services"
	"completeUser/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)

	uRepo := ur.New(db)
	uService := us.New(uRepo)
	ud.New(e, uService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":8000"))
}
