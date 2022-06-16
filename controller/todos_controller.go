package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"net/http"
	"todoservice/errhandler"
	"todoservice/services"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todos, err := services.GetAllTodos()
	if err != nil {
		errhandler.WriteError(r, w, errors.Wrap(err, ""), "GET /api/todos")
		return
	}

	jsonResponse, err := json.Marshal(todos)
	if err != nil {
		errhandler.WriteError(r, w, errors.Wrap(err, ""), "GET /api/todos")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todos, err := services.GetTodo(chi.URLParam(r, "id"))
	if err != nil {
		errhandler.WriteError(r, w, errors.Wrap(err, ""), "GET /api/todos/{id}")
		return
	}

	jsonResponse, err := json.Marshal(todos)
	if err != nil {
		errhandler.WriteError(r, w, errors.Wrap(err, ""), "GET /api/todo/{id}")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
