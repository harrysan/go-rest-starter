package config

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

// Load Config
var RootConfig AppConfig

type AppConfig struct {
	Host             string         `yaml:"host"`
	Port             string         `yaml:"port"`
	AppTimeoutConfig TimeoutConfig  `yaml:"timeout"`
	DatabaseConfig   DatabaseConfig `yaml:"database"`
	UserConfig       UserConfig     `yaml:"user"`
}

type TimeoutConfig struct {
	Server int `yaml:"server"`
	Read   int `yaml:"read"`
	Write  int `yaml:"write"`
	Idle   int `yaml:"idle"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type UserConfig struct {
	DefaultLoginPwd string `yaml:"defaultloginpwd"`
}

// Load
func init() {
	LoadConfigs()
}

func LoadConfigs() {
	yamlData, err := os.ReadFile("../../configs/config.yaml")

	if err != nil {
		log.Fatal("Error while reading app config file", err)
	}

	if err := yaml.Unmarshal(yamlData, &RootConfig); err != nil {
		panic(err)
	}
}
