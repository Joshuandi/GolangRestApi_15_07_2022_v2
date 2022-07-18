package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterRespone struct {
	R_username string `json:"username"`
	R_email    string `json:"email"`
}

type UserGetAll struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
