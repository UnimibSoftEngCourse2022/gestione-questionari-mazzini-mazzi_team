package main

import (
	common "form_management/common/logger"
	"form_management/internal"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		common.Logger.Fatal().Msg("||=> Error loading .env file")
	}
	internal.Server()
}
