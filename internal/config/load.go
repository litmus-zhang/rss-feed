package config

func NewConfig() (*Config, error) {

	config := Config{
		HttpServerAddress: "0.0.0.0:8080",
		DbDriver:          "postgres",
		DbSource:          "postgres://root:root@localhost:5432/rss_feed?sslmode=disable",
	}

	return &config, nil
}
