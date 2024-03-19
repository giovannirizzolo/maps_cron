package env

import (
	"github.com/joho/godotenv"
	"os"
)

func GetApiKey() string {
	err := LoadEnvModule()
	if  err != nil {
		return ""
	}
	
	apiKey := os.Getenv("MAPS_API_KEY")
	return apiKey
}

func LoadEnvModule() error {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	return nil
}