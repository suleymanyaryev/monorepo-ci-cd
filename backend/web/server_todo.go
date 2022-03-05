package web

import (
	"net/http"

	"example.com/monorepo-backend/responses"
)

func (s *Server) HandleGetToDoList(w http.ResponseWriter, r *http.Request) {
	err := responses.ErrOK
	responses.SendResponse(w, err, &TODOList)
}

func (s *Server) HandleCreateToDo(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	if len(name) == 0 {
		responses.SendResponse(w, responses.ErrBadRequest, nil)
		return
	}

	hasFound := false
	for _, todo := range TODOList {
		if todo.Name == name {
			hasFound = true
		}
	}

	if hasFound {
		responses.SendResponse(w, responses.ErrConflict, nil)
		return
	}

	todo := TODO{
		Name:   name,
		Status: "undone",
	}

	TODOList = append(TODOList, todo)

	err := responses.ErrOK
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

	hasFound := false
	tdList := make([]TODO, 0)

	for _, todo := range TODOList {
		if todo.Name != oldName {
			tdList = append(tdList, todo)
		} else {
			hasFound = true
			tdList = append(tdList, TODO{
				Name:   newName,
				Status: todo.Status,
			})
		}
	}

	if !hasFound {
		responses.SendResponse(w, responses.ErrNotFound, nil)
		return
	}

	TODOList = tdList
	responses.SendResponse(w, responses.ErrOK, nil)
}

func (s *Server) HandleDeleteToDo(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	if len(name) == 0 {
		responses.SendResponse(w, responses.ErrBadRequest, nil)
		return
	}

	hasFound := false
	tdList := make([]TODO, 0)
	for _, todo := range TODOList {
		if todo.Name != name {
			tdList = append(tdList, todo)
		} else {
			hasFound = true
		}
	}

	if !hasFound {
		responses.SendResponse(w, responses.ErrNotFound, nil)
		return
	}

	TODOList = tdList
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

	hasFound := false
	tdList := make([]TODO, 0)
	for _, todo := range TODOList {
		if todo.Name != name {
			tdList = append(tdList, todo)
		} else {
			hasFound = true
			tdList = append(tdList, TODO{
				Name:   todo.Name,
				Status: status,
			})
		}
	}

	if !hasFound {
		responses.SendResponse(w, responses.ErrNotFound, nil)
		return
	}

	TODOList = tdList
	responses.SendResponse(w, responses.ErrOK, nil)
}
