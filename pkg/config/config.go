package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type App struct {
	AppConfig        AppConfig      `mapstructure:"AppConfig"`
	AppTimeoutConfig TimeoutConfig  `mapstructure:"AppTimeout"`
	DatabaseConfig   DatabaseConfig `mapstructure:"AppDatabase"`
	RootConfig       RootConfig     `mapstructure:"AppRoot"`
}

type AppConfig struct {
	Host string
	Port string
}

type TimeoutConfig struct {
	Server int
	Read   int
	Write  int
	Idle   int
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

type RootConfig struct {
	Username string
	Password string
}

func LoadConfigs() App {
	var cfg App

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("Environment can't be loaded: ", err))
	}

	// fmt.Print("Host = ")
	// fmt.Println(viper.Get("App.Host"))
	// fmt.Print("Port = ")
	// fmt.Println(viper.Get("App.Port"))

	return cfg
}
