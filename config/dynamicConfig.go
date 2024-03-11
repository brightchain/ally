package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var GlobalConfig *viper.Viper

func init() {
	GlobalConfig = initViper()
	go dynamicReloadConfig()
}

func initViper() *viper.Viper {
	env := os.Getenv("GIN_MODE")
	name := "config"
	if env != "" {
		name += "_" + env
	}
	GlobalConfig = viper.New()
	GlobalConfig.SetConfigName(name)
	GlobalConfig.SetConfigType("yaml")
	GlobalConfig.AddConfigPath(".")

	err := GlobalConfig.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	return GlobalConfig
}

func dynamicReloadConfig() {
	GlobalConfig.WatchConfig()
}
