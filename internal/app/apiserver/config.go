package apiserver

type Config struct {
	BindAddr string `toml:"bin_addr"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
	}
}