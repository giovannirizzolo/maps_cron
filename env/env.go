package env

import (
	"os"
)

func GetApiKey() string {
	apiKey := os.Getenv("MAPS_API_KEY")
	return apiKey
}