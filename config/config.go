package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Port int    `mapstructure:"port"`
		Name string `mapstructure:"name"`
		Env  string `mapstructure:"env"`
	} `mapstructure:"app"`

	Mongo struct {
		URI      string `mapstructure:"uri"`
		Database string `mapstructure:"database"`
	} `mapstructure:"mongo"`

	JWT struct {
		Secret    string `mapstructure:"secret"`
		ExpiresIn string `mapstructure:"expires_in"`
	} `mapstructure:"jwt"`

	Logging struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
	} `mapstructure:"logging"`
}

var AppConfig *Config

func LoadConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	viper.SetConfigName("config." + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	AppConfig = &cfg
	return &cfg, nil
}

func Load() {
	_, err := LoadConfig()
	if err != nil {
		panic(err)
	}
}
