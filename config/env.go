package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAdress               string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error to load environment variables file. ", err)
	}

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "localhost"),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "mypassword123!"),
		DBAdress:               fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3307")),
		DBName:                 getEnv("DB_NAME", "ecom"),
		JWTExpirationInSeconds: getEnvAsInt("JWT,EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-anymore"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
