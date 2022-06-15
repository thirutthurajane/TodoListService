package controller

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"todoservice/errhandler"
	"todoservice/services"
)

func GithubAuthCallback(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	provider := "github"

	result, err := services.HandleCallback(provider, code)
	if err != nil {
		errhandler.WriteError(r, w, errors.Wrap(err, ""), "GET /auth/github/callback")
		return
	}
	jsonResponse, _ := json.Marshal(result)

	w.Write(jsonResponse)
}
