package web

import (
	"net/http"
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
