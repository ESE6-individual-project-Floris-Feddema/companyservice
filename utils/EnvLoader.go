package utils

import (
	"github.com/joho/godotenv"
	"os"
)

func EnvVar(key string) string {
	godotenv.Load(".env")
	return os.Getenv(key)
}
