package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ListenAddress string `json:"listen_address"`
	DbConn        string `json:"db_conn"`
}

var Conf *Config

func ReadConfig(source string) (err error) {
	if _, err := os.Stat(source); !(errors.Is(err, os.ErrNotExist)) {
		err1 := godotenv.Load(source)
		if err1 != nil {
			return err1
		}
	}

	Conf = &Config{
		ListenAddress: os.Getenv("LISTEN_ADDR"),
		DbConn: fmt.Sprintf(
			"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
		),
	}

	return
}
