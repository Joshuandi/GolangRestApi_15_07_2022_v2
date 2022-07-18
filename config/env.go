package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	MongoDb string `env:"MongoDbConnect"`
	PORT    int    `env:"PORT"`
	Mysql   string `env:"MysqlConnect"`
}

func EnvMysql() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MysqlConnect")
}
