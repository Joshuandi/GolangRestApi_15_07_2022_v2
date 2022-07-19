package config

import (
	"GolangRestApi_15_07_2022_v2/model"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Redis *redis.Client
	Gdb   *gorm.DB
}

var cfg EnvConfig

func ConnectMysqlGorm() *DB {
	_ = cleanenv.ReadConfig(".env", &cfg)
	redisClient := RedisClient()
	ping, err := PingRedis(redisClient)
	if err != nil {
		fmt.Println("error connect redis:", err.Error())
	} else {
		fmt.Println("PING:", ping)
	}
	Gdb, err := gorm.Open(mysql.Open(cfg.Mysql), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Database Successfuly Connected PORT :", cfg.PORT)
	Gdb.AutoMigrate(&model.Users{})
	return &DB{Redis: redisClient, Gdb: Gdb}
}
func ConnectMysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", EnvMysql())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Successfuly Connected PORT :")
	return db, nil
}

func RedisClient() *redis.Client {
	rC := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rC
}

func PingRedis(client *redis.Client) (string, error) {
	ping, err := client.Ping().Result()
	if err != nil {
		return "Error Ping", err
	} else {
		return ping, nil
	}
}

var Db *DB = ConnectMysqlGorm()
