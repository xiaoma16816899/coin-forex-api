package utils

import (
	"fmt"
	"os"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) string {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "mysql":
		// URL for Mysql connection.
		url = fmt.Sprintf(
			"%s:%s@tcp(%s)/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
		)
	case "redis":
		url = fmt.Sprintf(
			"redis://%s:%s@%s:%s/%s",
			os.Getenv("REDIS_USER"),
			os.Getenv("REDIS_PASSWORD"),
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
			os.Getenv("REDIS_DB_NUMBER"),
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		// Return error message.
		return ""
	}

	// Return connection URL.
	return url
}
