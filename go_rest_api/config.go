package main

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mysql123"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "mysql_docker"), getEnv("DB_PORT", "33060")),
		DBName:     getEnv("DB_NAME", "projectmanager"),
		JWTSecret:  getEnv("JWT_SECRET", "randomjwtsecretkey"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
