package config

import (
	"os"

	"github.com/go-sql-driver/mysql"
)

func NewDBConfig() mysql.Config {
	return mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_ADDR"),
		DBName: os.Getenv("DB_NAME"),
	}
}
