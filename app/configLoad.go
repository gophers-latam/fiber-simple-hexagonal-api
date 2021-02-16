package app

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	logs "github.com/zeroidentidad/fiber-hex-api/logger"
)

func configLoad() {
	err := godotenv.Load()
	if err != nil {
		logs.Fatal("Error loading .env file")
	}

	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logs.Fatal(fmt.Sprintf("- %s not defined. Terminating app...", k))
		}
	}
}
