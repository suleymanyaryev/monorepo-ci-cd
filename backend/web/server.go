package web

import (
	"github.com/gorilla/mux"
)

type Server struct{}

type TODO struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

var TODOList []TODO = make([]TODO, 0)

func NewServer() *Server {
	return &Server{}
}

func NewRouter(s *Server) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/todo/list", s.HandleGetToDoList).Methods("GET")
	r.HandleFunc("/api/v1/todo/create", s.HandleCreateToDo).Methods("POST")
	r.HandleFunc("/api/v1/todo/update", s.HandleUpdateToDo).Methods("POST")
	r.HandleFunc("/api/v1/todo/change-status", s.HandleChangeToDoState).Methods("POST")
	r.HandleFunc("/api/v1/todo/delete", s.HandleDeleteToDo).Methods("POST")
	return r
}
