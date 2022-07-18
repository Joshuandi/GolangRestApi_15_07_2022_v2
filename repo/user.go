package repo

import (
	"GolangRestApi_15_07_2022_v2/config"
	"GolangRestApi_15_07_2022_v2/model"
	"fmt"

	"gorm.io/gorm"
)

type UserRepoInterface interface {
	UserRepoRegister(users model.Users) (*model.Users, error)
	UserRepoGetAll() (*[]model.Users, error)
	UserRepoFindById(users model.Users, id string) (*model.Users, error)
	UserRepoPut(users model.Users, id string) (*model.Users, error)
	UserRepoDelete(users model.Users, id string) (*model.Users, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &UserRepo{db: db}
}

func (u UserRepo) UserRepoRegister(users model.Users) (*model.Users, error) {
	fmt.Println("parsing data create:", users)
	err := config.Gdb.Debug().Create(&users).Error
	fmt.Println("hasil create:", users)
	if err != nil {
		fmt.Println("Create row error :", err)
	}
	return &users, nil
}

func (u UserRepo) UserRepoGetAll() (*[]model.Users, error) {
	var users []model.Users
	err := config.Gdb.Debug().Find(&users).Error
	fmt.Println("hasil get all:", users)
	if err != nil {
		fmt.Println("Find All row error :", err)
	}
	return &users, nil
}

func (u UserRepo) UserRepoFindById(users model.Users, id string) (*model.Users, error) {
	err := config.Gdb.Debug().First(users, id)
	fmt.Println("hasil get by id:", users)
	if err != nil {
		fmt.Println("Find by id row error :", err)
	}
	return &users, nil
}

func (u UserRepo) UserRepoPut(users model.Users, id string) (*model.Users, error) {
	err := config.Gdb.Debug().First(users, id)
	fmt.Println("hasil update by id:", users)
	if err != nil {
		fmt.Println("Find by id row error :", err)
	}
	err2 := u.db.Save(&users)
	fmt.Println("hasil update by id save:", users)
	if err2 != nil {
		fmt.Println("Save/Put by id row error :", err2)
	}
	return &users, nil
}

func (u UserRepo) UserRepoDelete(users model.Users, id string) (*model.Users, error) {
	err := config.Gdb.Debug().Delete(users, id)
	fmt.Println("hasil delete by id:", users)
	if err != nil {
		fmt.Println("Delete by id row error :", err)
	}
	return &users, nil
}
