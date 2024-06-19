package main

import (
	"github.com/joho/godotenv"
	"github.com/nlypage/smart-home-prefix-project/cmd/app"
	"github.com/nlypage/smart-home-prefix-project/internal/adapters/config"
	"github.com/nlypage/smart-home-prefix-project/internal/adapters/controller/api/setup"
	"log"
)

// main is entry point of the program.
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	appConfig := config.GetConfig()
	smartHomeApp := app.NewSmartHomeApp(appConfig)

	setup.Setup(smartHomeApp)
	smartHomeApp.Start()
}
