package service

import (
	"GolangRestApi_15_07_2022_v2/model"
	"GolangRestApi_15_07_2022_v2/repo"
	"GolangRestApi_15_07_2022_v2/util"
	"errors"
	"fmt"
)

type UserServiceInterface interface {
	UserServiceRegister(users model.Users) (*model.Users, error)
	UserServiceGetAll() (*[]model.Users, error)
	UserServiceGetById(users model.Users, id string) (*model.Users, error)
	UserServicePut(users model.Users, id string) (*model.Users, error)
	UserServiceDelete(users model.Users, id string) (*model.Users, error)
}

type UserService struct {
	userRepo repo.UserRepoInterface
}

func NewUserService(userRepo repo.UserRepoInterface) UserServiceInterface {
	return &UserService{userRepo: userRepo}
}

func (u UserService) UserServiceRegister(users model.Users) (*model.Users, error) {
	email := users.Email
	username := users.Username
	if _, ok := util.ValidateEmail(users.Email); !ok {
		return nil, errors.New("Email must valid")
	}
	if email == "" {
		return nil, errors.New("Email must be input")
	}
	if username == "" {
		return nil, errors.New("Username must be input")
	}
	if users.Password == "" {
		return nil, errors.New("Password must be input")
	}
	if len(users.Password) < 6 {
		return nil, errors.New("Password must more than 6 character")
	}
	pass, errHash := util.GenerateHashPassword(users.Password)
	if errHash != nil {
		fmt.Println("Error Hash : " + errHash.Error())
		return nil, errHash
	}
	users.Password = pass
	fmt.Println("ini service users:", users)

	userRegis, err := u.userRepo.UserRepoRegister(users)

	if err != nil {
		fmt.Println("Error While Register", err.Error())
		return nil, err
	}
	return userRegis, nil
}

func (u UserService) UserServiceGetAll() (*[]model.Users, error) {
	user, err := u.userRepo.UserRepoGetAll()
	fmt.Println("get all:", user)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return user, nil
}

func (u UserService) UserServiceGetById(users model.Users, id string) (*model.Users, error) {
	user, err := u.userRepo.UserRepoFindById(users, id)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return user, nil
}

func (u UserService) UserServicePut(users model.Users, id string) (*model.Users, error) {
	email := users.Email
	username := users.Username
	if _, ok := util.ValidateEmail(email); !ok {
		return nil, errors.New("Email must valid")
	}
	if email == "" {
		return nil, errors.New("Email must be input")
	}
	if username == "" {
		return nil, errors.New("Username must be input")
	}
	user, err := u.userRepo.UserRepoPut(users, id)
	if err != nil {
		fmt.Println("Error While Update", err.Error())
		return nil, err
	}
	return user, nil
}

func (u UserService) UserServiceDelete(users model.Users, id string) (*model.Users, error) {
	deleteUser, err := u.userRepo.UserRepoDelete(users, id)
	if err != nil {
		return nil, err
	}
	return deleteUser, nil
}
