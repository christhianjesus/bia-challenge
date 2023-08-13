package config

type Config struct {
	SRVHost string `env:"HOST,default=0.0.0.0"`
	SRVPort string `env:"PORT,default=8080"`

	DBConfig
}

func (c Config) Addr() string {
	return c.SRVHost + ":" + c.SRVPort
}
