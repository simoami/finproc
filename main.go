package main

import (
  "github.com/spf13/viper"

  "finproc/app/models"
  "finproc/app"
  "finproc/app/config"
)

const local string = "LOCAL"

func main() {

  // apply environment and config file
  config.Init()

	// initialse application context
	ctx := models.AppContext{
		// Version: version,
		Env:     viper.GetString("APP_ENV"),
		Port:    viper.GetInt("PORT"),
	}

	// start application
	app.StartServer(ctx)
}
