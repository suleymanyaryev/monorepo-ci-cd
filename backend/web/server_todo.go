package web

import (
	"net/http"

	"example.com/monorepo-backend/responses"
)

func (s *Server) HandleGetToDoList(w http.ResponseWriter, r *http.Request) {
	data, err := s.pg.GetToDoList(r.Context())
	if err != nil {
		responses.SendResponse(w, responses.ErrInternalServerError, nil)
		return
	}
	err = responses.ErrOK
	responses.SendResponse(w, err, data)
}

func (s *Server) HandleCreateToDo(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	if len(name) == 0 {
		responses.SendResponse(w, responses.ErrBadRequest, nil)
		return
	}

	toDoWithName, err := s.pg.GetToDoByName(r.Context(), name)
	if err != nil {
		responses.SendResponse(w, responses.ErrInternalServerError, nil)
		return
	}
	if toDoWithName != nil {
		responses.SendResponse(w, responses.ErrConflict, nil)
		return
	}

	todo := responses.ToDO{
		Name:   name,
		Status: "undone",
	}

	err = s.pg.CreateToDo(r.Context(), &todo)
	if err != nil {
		responses.SendResponse(w, responses.ErrInternalServerError, nil)
		return
	}

	err = responses.ErrOK
	responses.SendResponse(w, err, &todo)
}

func (s *Server) HandleUpdateToDo(w http.ResponseWriter, r *http.Request) {

	oldName := r.FormValue("old_name")
	if len(oldName) == 0 {
		responses.SendResponse(w, responses.ErrBadRequest, nil)
		return
	}

	newName := r.FormValue("new_name")
	if len(newName) == 0 {
		responses.SendResponse(w, responses.ErrBadRequest, nil)
		return
	}

	if newName == oldName {
		responses.SendResponse(w, responses.ErrOK, nil)
		return
	}

	oldTodo, err := s.pg.GetToDoByName(r.Context(), oldName)
	if err != nil {
		responses.SendResponse(w, responses.ErrInternalServerError, nil)
		return
	}
	if oldTodo == nil {
		responses.SendResponse(w, responses.ErrNotFound, nil)
		return
	}

	toDoWithNewName, err := s.pg.GetToDoByName(r.Context(), newName)
	if err != nil {
		responses.SendResponse(w, responses.ErrInternalServerError, nil)
		return
	}
	if toDoWithNewName != nil {
		responses.SendResponse(w, responses.ErrConflict, nil)
		return
	}

	err = s.pg.UpdateToDo(r.Context(), oldName, newName)
	if err != nil {
		responses.SendResponse(w, responses.ErrInternalServerError, nil)
		return
	}

	responses.SendResponse(w, responses.ErrOK, nil)
}

func (s *Server) HandleDeleteToDo(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	if len(name) == 0 {
		responses.SendResponse(w, responses.ErrBadRequest, nil)
		return
	}

	todo, err := s.pg.GetToDoByName(r.Context(), name)
	if err != nil {
		responses.SendResponse(w, responses.ErrInternalServerError, nil)
		return
	}
	if todo == nil {
		responses.SendResponse(w, responses.ErrNotFound, nil)
		return
	}

	err = s.pg.DeleteToDo(r.Context(), name)
	if err != nil {
		responses.SendResponse(w, responses.ErrInternalServerError, nil)
		return
	}

	responses.SendResponse(w, responses.ErrOK, nil)
}

func (s *Server) HandleChangeToDoState(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	if len(name) == 0 {
		responses.SendResponse(w, responses.ErrBadRequest, nil)
		return
	}

	status := r.FormValue("status")
	if len(status) == 0 {
		responses.SendResponse(w, responses.ErrBadRequest, nil)
		return
	}

	todo, err := s.pg.GetToDoByName(r.Context(), name)
	if err != nil {
		responses.SendResponse(w, responses.ErrInternalServerError, nil)
		return
	}
	if todo == nil {
		responses.SendResponse(w, responses.ErrNotFound, nil)
		return
	}

	err = s.pg.ChangeToDoStatus(r.Context(), name, status)
	if err != nil {
		responses.SendResponse(w, responses.ErrInternalServerError, nil)
		return
	}

	responses.SendResponse(w, responses.ErrOK, nil)
}
