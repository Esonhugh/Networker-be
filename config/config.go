package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	*viper.Viper
}

// GlobalConfig 默认全局配置
var GlobalConfig *Config

// Init 使用 ./application.yaml 初始化全局配置
func Init() {
	GlobalConfig = &Config{
		viper.New(),
	}
	GlobalConfig.SetConfigName("application")
	GlobalConfig.SetConfigType("yaml")
	GlobalConfig.AddConfigPath(".")
	GlobalConfig.AddConfigPath("../") // For Debug
	GlobalConfig.AddConfigPath("./config")

	err := GlobalConfig.ReadInConfig()
	if err != nil {
		log.Panic("Config Reading Error", err)
	}
}
