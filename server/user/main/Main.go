package main

import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	var env = "development"

	viper.AddConfigPath("../config")
	viper.SetConfigName("config-" + env)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

}
