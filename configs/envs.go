package configs

import (
	"fmt"
	"os"
	"strconv"
)

type Env struct {
	APP_PORT    int
	DB_DATABASE string
	DB_USERNAME string
	DB_PASSWORD string
	DB_SSLMODE  string
	DB_HOST     string
	DB_PORT     int
}

func GetEnvs() *Env {
	return &Env{
		DB_DATABASE: MustGetString("DB_DATABASE"),
		DB_USERNAME: MustGetString("DB_USERNAME"),
		DB_PASSWORD: MustGetString("DB_PASSWORD"),
		DB_SSLMODE:  MustGetString("DB_SSLMODE"),
		DB_HOST:     MustGetString("DB_HOST"),
		APP_PORT:    MustGetInt("APP_PORT"),
		DB_PORT:     MustGetInt("DB_PORT"),
	}
}

func MustGetString(name string) string {
	value := os.Getenv(name)
	if value == "" {
		fmt.Println("%s can't be empty\n", name)
		os.Exit(1)
	}

	return value
}

func MustGetInt(name string) int {
	value, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		fmt.Println("%s must contain a float value!\n", name)
		os.Exit(1)
	}

	return value
}
