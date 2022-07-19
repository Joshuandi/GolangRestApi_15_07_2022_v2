package repo

import (
	"GolangRestApi_15_07_2022_v2/config"
	"GolangRestApi_15_07_2022_v2/model"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var conn *gorm.DB = config.Db.Gdb
var r *redis.Client = config.Db.Redis

type UserRepoInterface interface {
	UserRepoRegister(e echo.Context, users model.Users) (*model.Users, error)
	UserRepoGetAll(e echo.Context) (*[]model.Users, error)
	UserRepoFindById(e echo.Context, users model.Users, id string) (*model.Users, error)
	UserRepoPut(e echo.Context, users model.Users, id string) (*model.Users, error)
	UserRepoDelete(users model.Users, id string) (*model.Users, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &UserRepo{db: db}
}

func (u UserRepo) UserRepoRegister(e echo.Context, users model.Users) (*model.Users, error) {
	//fmt.Println("parsing data create:", users)
	err := conn.Debug().Create(&users).Error
	//fmt.Println("hasil create:", users)
	if err != nil {
		fmt.Println("CREATE row error :", err)
		return &users, err
	}
	return &users, nil
}

func (u UserRepo) UserRepoGetAll(e echo.Context) (*[]model.Users, error) {
	var users []model.Users
	err := conn.Debug().Find(&users).Error
	//fmt.Println("hasil get all:", users)
	if err != nil {
		fmt.Println("Find All row error :", err)
		return &users, err
	}
	return &users, nil
}

func (u UserRepo) UserRepoFindById(e echo.Context, users model.Users, id string) (*model.Users, error) {
	//var valueRedis *model.Users = cache.RedisCacheInterface.Get(id)
	err := conn.Debug().First(&users, id).Error
	//fmt.Println("hasil get by id:", users)
	if err != nil {
		fmt.Println("Find by id row error :", err)
		return &users, err
	}
	return &users, nil
}

func (u UserRepo) UserRepoPut(e echo.Context, users model.Users, id string) (*model.Users, error) {
	var modelUser model.Users
	err := conn.Debug().First(&modelUser, id).Error
	//fmt.Println("hasil update by id:", users)
	if err != nil {
		fmt.Println("Find by id row error :", err)
		return &users, err
	}
	err2 := conn.Model(&modelUser).Updates(users).Error
	//fmt.Println("hasil update by id save:", users)
	if err2 != nil {
		fmt.Println("Save/Put by id row error :", err2)
		return &users, err
	}
	return &users, nil
}

func (u UserRepo) UserRepoDelete(users model.Users, id string) (*model.Users, error) {
	err := conn.Debug().Delete(users, id).Error
	//fmt.Println("hasil delete by id:", users)
	if err != nil {
		fmt.Println("Delete by id row error :", err)
		return &users, err
	}
	return &users, nil
}
