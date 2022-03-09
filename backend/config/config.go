package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ListenAddress string `json:"listen_address"`
}

var Conf *Config

func ReadConfig(source string) (err error) {
	os.Stat(".env")
	if _, err := os.Stat(".env"); !(errors.Is(err, os.ErrNotExist)) {
		err1 := godotenv.Load()
		if err1 != nil {
			return err1
		}
	}

	Conf = &Config{
		ListenAddress: os.Getenv("LISTEN_ADDR"),
	}

	return
}
