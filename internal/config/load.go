package config

func NewConfig() (*Config, error) {

	config := Config{
		HttpServerAddress: "0.0.0.0:8080",
		DbDriver:          "postgres",
		DbSource:          "rss_feed",
	}

	return &config, nil
}
