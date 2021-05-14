package config

type Config struct {
	Debug  bool
	Port   string
	DBPath string
}

func NewConfig() *Config {
	return &Config{
		Debug:  true,
		Port:   ":8080",
		DBPath: "./example.db",
	}
}
