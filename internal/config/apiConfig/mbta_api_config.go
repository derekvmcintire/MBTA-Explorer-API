package apiConfig

import "os"

func GetAPIKey() string {
	return os.Getenv("MBTA_API_KEY")
}
