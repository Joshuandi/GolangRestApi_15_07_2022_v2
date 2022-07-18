package config

import (
	"GolangRestApi_15_07_2022_v2/model"
	"database/sql"
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Gdb *gorm.DB
var cfg EnvConfig

func ConnectMysqlGorm() (*gorm.DB, error) {
	_ = cleanenv.ReadConfig(".env", &cfg)
	Gdb, err := gorm.Open(mysql.Open(cfg.Mysql), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Database Successfuly Connected PORT :", cfg.PORT)
	Gdb.AutoMigrate(&model.Users{})
	return Gdb, nil
}

func ConnectMysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", EnvMysql())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Successfuly Connected PORT :")
	return db, nil
}
