package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	AppID      string
	AppSecret  string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPass     string
	DbName     string
	JwtSecret  string
	JwtTTL     int
}

func LoadConfig() *Config {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	serverPort := os.Getenv("SERVER_PORT")
	appID := os.Getenv("APP_ID")
	appSecret := os.Getenv("APP_SECRET")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtTTL, _ := strconv.Atoi(os.Getenv("JWT_TTL"))

	return &Config{
		ServerPort: serverPort,
		AppID:      appID,
		AppSecret:  appSecret,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbUser:     dbUser,
		DbPass:     dbPass,
		DbName:     dbName,
		JwtSecret:  jwtSecret,
		JwtTTL:     jwtTTL,
	}
}
