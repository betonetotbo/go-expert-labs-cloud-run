package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port          int    `mapstructure:"PORT"`
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
}

func Load() (*Config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	viper.SetDefault("PORT", 8080)
	viper.SetDefault("WEATHER_API_KEY", nil)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
