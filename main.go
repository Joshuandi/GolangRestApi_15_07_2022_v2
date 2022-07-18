package main

import (
	"fmt"

	"GolangRestApi_15_07_2022_v2/config"
	"GolangRestApi_15_07_2022_v2/router"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	_, err := config.Db.DB()
	if err != nil {
		fmt.Println("error connect:", err.Error())
		return
	}
	//router
	router.UserRouter(e)

	e.Logger.Fatal(e.Start(":8088"))
}
