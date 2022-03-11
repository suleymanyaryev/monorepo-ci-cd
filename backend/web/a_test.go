package web

import (
	"net/http"

	"example.com/monorepo-backend/config"
	log "github.com/sirupsen/logrus"
)

type Case struct {
	name     string
	r        *http.Request
	wantCode int
	wantData map[string]interface{}
}

type ListCase struct {
	name     string
	r        *http.Request
	wantCode int
	wantData []interface{}
}

//

func init() {
	err := config.ReadConfig("../.env")
	if err != nil {
		log.WithError(err).Error("error in reading env")
	}
}
