package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ListenAddress string `json:"listen_address"`
}

var Conf *Config

func ReadConfig(source string) (err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}

	Conf = &Config{
		ListenAddress: os.Getenv("LISTEN_ADDR"),
	}

	return
}
