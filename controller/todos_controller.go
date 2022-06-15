package controller

import (
	"encoding/json"
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
