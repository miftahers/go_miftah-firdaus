package config

import (
	"os"
)

type Config struct {
	APIPort     string
	APIKey      string
	TokenSecret string
}

var Cfg *Config

func InitConfig() {
	cfg := &Config{}

	// viper.SetConfigName("app")
	// viper.SetConfigFile("env")
	// viper.AddConfigPath(".")
	// viper.ReadInConfig()
	// viper.Unmarshal(&cfg)

	cfg.APIPort = SetEnv("APIPort", ":8080")
	cfg.APIKey = SetEnv("APIKey", "AppSuberb-WAW")
	cfg.TokenSecret = "SuperbMIFTAH"

	Cfg = cfg
}

func SetEnv(key, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}
