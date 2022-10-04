package apiserver

type Config struct {
	BindAddr string `toml:"bin_addr"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":9001",
		LogLevel: "debug",
	}
}