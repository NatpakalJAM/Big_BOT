package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//Config main config
type Config struct {
	Token string `mapstructure:"token"`
}

//C config instance
var C Config

//Init init config
func Init() {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	err = viper.Unmarshal(&C)
	if err != nil {
		panic(err)
	}
}
