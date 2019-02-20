package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

//Config representa configuração de conexao com banco
type Config struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
