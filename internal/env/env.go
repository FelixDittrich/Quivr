package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnvFile(envPath string) {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func GetString(key string, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}