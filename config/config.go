package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	PORT        string
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
	DB_HOST     string
	DB_PORT     string
}

var ENV Config

// LoadConfig reads the configuration from the .env file and unmarshals it
// into the ENV Config struct. If the configuration file cannot be read or
// unmarshaled, the function logs a fatal error message. Otherwise, the function
// logs a success message to the console.
func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal(err)
	}

	log.Println("Load server is successfully...")
}
