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
	JWTConfig        JWTConfig      `mapstructure:"JWTConfig"`
	RedisConfig      RedisConfig    `mapstructure:"RedisConfig"`
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

type JWTConfig struct {
	JWTSecretKey string
	TokenExpiry  int // Expiry in minutes
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func LoadConfigs() App {
	var cfg App

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	// Handle errors reading the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("environment can't be loaded: %w", err))
	}

	return cfg
}
