package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

// mysql 配置读取 .env
var Mysql = map[string]string{
	"host":     os.Getenv("DB_HOST"),
	"port":     os.Getenv("DB_PORT"),
	"database": os.Getenv("DB_DATABASE"),
	"username": os.Getenv("DB_USERNAME"),
	"password": os.Getenv("DB_PASSWORD"),
}
