package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading the .env file!\nError: %v", err)
	}
}

func GetEnv(k string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading the .env file!\nError: %v", err)
	}

	v, there := os.LookupEnv(k)
	if !there {
		return "", fmt.Errorf("Cannot find the env variable - %s", k)
	}

	return v, nil
}
