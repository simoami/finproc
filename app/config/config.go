package config

import (
	"log"

	"github.com/gin-gonic/gin"
  "github.com/spf13/viper"
)

// Init config settings from the env, and config file(s
func Init() {
  viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./app/config")

  setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("config file not found, using defaults")
	}

	gin.SetMode(viper.GetString("GIN_MODE"))
}

// Default config settings if not found in the env or config file(s)
func setDefaults() {
  viper.SetDefault("APP_ENV", "local") // local, dev, stage, prod
  viper.SetDefault("APP_DEBUG", false)
  viper.SetDefault("GIN_MODE", "release") // debug or release
  viper.SetDefault("PORT", 8080)
}
