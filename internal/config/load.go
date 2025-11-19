package config

import "os"

func NewConfig() (*Config, error) {

	config := Config{
		HttpServerAddress: "0.0.0.0:8080",
		DbDriver:          os.Getenv("DB_DRIVER"),
		DbSource:          os.Getenv("DB_SOURCE"),
	}

	return &config, nil
}
