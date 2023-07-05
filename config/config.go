package config

import (
	"os"
	"strconv"
)

type Config struct {
	DB *DBConfig
	JWTSecret string
	Host string
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     os.Getenv("DB_HOST"),
			Port:     port,
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Charset:  "utf8",
		},
		JWTSecret: os.Getenv("JWT_SECRET"),
		Host: os.Getenv("HOST"),
	}
}
