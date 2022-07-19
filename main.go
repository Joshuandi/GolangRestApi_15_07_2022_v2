package main

import (
	"fmt"

	"GolangRestApi_15_07_2022_v2/config"
	"GolangRestApi_15_07_2022_v2/handler"
	"GolangRestApi_15_07_2022_v2/repo"
	"GolangRestApi_15_07_2022_v2/router"
	"GolangRestApi_15_07_2022_v2/service"

	"github.com/labstack/echo/v4"
)

func main() {
	//echo
	e := echo.New()
	//connect database
	_, err := config.Db.Gdb.DB()
	if err != nil {
		fmt.Println("error connect:", err.Error())
		return
	}
	//interface user 
	usersRepo := repo.NewUserRepo(config.Db.Gdb)
	userService := service.NewUserService(usersRepo)
	userHandler := handler.NewUserHandler(userService)
	//router
	router.UserRouter(e, userHandler)

	e.Logger.Fatal(e.Start(":8088"))
}
