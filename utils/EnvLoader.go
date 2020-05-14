package utils

import (
	"github.com/joho/godotenv"
	"os"
)

func EnvVar(key string) string {
	godotenv.Load("config.env")
	return os.Getenv(key)
}
