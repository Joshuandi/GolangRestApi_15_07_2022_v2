package router

import (
	"GolangRestApi_15_07_2022_v2/handler"

	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Echo, u *handler.UserHandler) {
	e.POST("/users/register", u.UserHandlerRegister)
	e.GET("/users", u.UserHandlerGetAll)
	e.GET("/users/:id", u.UserHandlerGetById)
	e.PUT("/users/:id", u.UserHandlerPut)
	e.DELETE("/users/:id", u.UserHandlerDelete)
}
