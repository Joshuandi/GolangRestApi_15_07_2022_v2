package handler

import (
	"GolangRestApi_15_07_2022_v2/model"
	"GolangRestApi_15_07_2022_v2/response"
	"GolangRestApi_15_07_2022_v2/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) UserHandlerRegister(e echo.Context) error {
	var users model.Users
	if err := e.Bind(&users); err != nil {
		return e.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &echo.Map{"data": err.Error()}})
	}
	fmt.Println("ini bind register user:", users)
	result, err := u.userService.UserServiceRegister(e, users)
	if err != nil {
		return e.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &echo.Map{"data": err.Error()}})
	}
	fmt.Println("register:", users)
	return e.JSON(http.StatusCreated,
		response.UserResponse{
			Status:  http.StatusCreated,
			Message: "Successfully Register",
			Data:    &echo.Map{"data": result}})
}

func (u *UserHandler) UserHandlerGetAll(e echo.Context) error {
	result, err := u.userService.UserServiceGetAll(e)
	fmt.Println("get all:", result)
	if err != nil {
		return e.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &echo.Map{"data": err.Error()}})
	}
	return e.JSON(http.StatusCreated,
		response.UserResponse{
			Status:  http.StatusCreated,
			Message: "User Data",
			Data:    &echo.Map{"data": &result}})
}

func (u *UserHandler) UserHandlerGetById(e echo.Context) error {
	id := e.Param("id")
	var users model.Users
	result, err := u.userService.UserServiceGetById(e, users, id)
	if err != nil {
		return e.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &echo.Map{"data": err.Error()}})
	}
	return e.JSON(http.StatusCreated,
		response.UserResponse{
			Status:  http.StatusCreated,
			Message: "User Data",
			Data:    &echo.Map{"data": &result}})
}

func (u *UserHandler) UserHandlerPut(e echo.Context) error {
	id := e.Param("id")
	var users model.Users
	if err := e.Bind(&users); err != nil {
		return e.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &echo.Map{"data": err.Error()}})
	}
	result, err := u.userService.UserServicePut(e, users, id)
	if err != nil {
		return e.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &echo.Map{"data": err.Error()}})
	}
	update := model.UserRegisterRespone{
		R_username: result.Username,
		R_email:    result.Email,
	}
	return e.JSON(http.StatusCreated,
		response.UserResponse{
			Status:  http.StatusCreated,
			Message: "Successfully Updated",
			Data:    &echo.Map{"data": update}})
}

func (u *UserHandler) UserHandlerDelete(e echo.Context) error {
	id := e.Param("id")
	if index, err := strconv.Atoi(id); err == nil {
		_, err := u.userService.UserServiceDelete(model.Users{Id: index}, id)
		if err != nil {
			return e.JSON(http.StatusBadRequest,
				response.UserResponse{
					Status:  http.StatusBadRequest,
					Message: "error",
					Data:    &echo.Map{"data": err.Error()}})
		}
	}
	return e.JSON(http.StatusCreated,
		response.UserResponse{
			Status:  http.StatusCreated,
			Message: "Successfully Deleted",
			Data:    &echo.Map{"data": "Deleted Data"}})
}
