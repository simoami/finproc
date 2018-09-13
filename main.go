package main

import (
	"github.com/spf13/viper"

	"finproc/app"
	"finproc/app/config"
	"finproc/app/models"
)

func main() {

	// apply environment and config file
	config.Init()

	// initialse application context
	ctx := models.AppContext{
		// Version: version,
		Env:  viper.GetString("APP_ENV"),
		Port: viper.GetInt("PORT"),
	}

	// start application
	app.StartServer(ctx)
}
