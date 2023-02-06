package config

import "os"

type Config struct {
	Token string
}

func (c *Config) ReadConfig() *Config {
	c.Token = os.Getenv("TOKEN")

	return c
}
