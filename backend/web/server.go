package web

import (
	"example.com/monorepo-backend/datastore"
	"github.com/gorilla/mux"
)

type Server struct {
	pg *datastore.PgAccess
}

func NewServer(d *datastore.PgAccess) *Server {
	return &Server{
		pg: d,
	}
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
