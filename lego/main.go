package main

import (
	"log"
	"other-side/cmd"
	"other-side/config"

	"github.com/joho/godotenv"
)

func main() {
	if godotenv.Load(".env") != nil {
		log.Fatal("Error loading .env file")
	}

	confVars, configErr := config.New()

	if configErr != nil {
		log.Fatal(configErr)
	}

	app := cmd.InitApp()

	err := app.Listen(confVars.Port)
	if err != nil {
		log.Fatal(err)
	}
}
