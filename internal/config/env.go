package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	PORT string
	DSN  string
}

func LoadEnv() *Env {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error: loading .env file")
	}

	return &Env{
		PORT: os.Getenv("PORT"),
		DSN:  os.Getenv("DSN"),
	}
}
