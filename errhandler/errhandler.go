package errhandler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type ErrResponse struct {
	TimeStamps time.Time `json:"timeStamps"`
	Status     string    `json:"status"`
	IsError    bool      `json:"isError"`
	Message    string    `json:"message"`
	ApiPath    string    `json:"apiPath"`
}

var logger = logrus.New()

func init() {
	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
	}
}

func WriteError(req *http.Request, res http.ResponseWriter, err error, path string) {
	response := ErrResponse{
		TimeStamps: time.Now(),
		Status:     "500",
		IsError:    true,
		Message:    err.Error(),
		ApiPath:    path,
	}
	WriteErrorLog(req, errors.Cause(err).Error())
	jsonResponse, _ := json.Marshal(response)
	res.WriteHeader(http.StatusInternalServerError)
	res.Write(jsonResponse)
	return
}

func WriteErrorLog(req *http.Request, err string) {
	reqID := middleware.GetReqID(req.Context())
	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}
	errLog := logger.WithFields(logrus.Fields{
		"ts":          time.Now().UTC().Format(time.RFC1123),
		"req_id":      reqID,
		"http_scheme": scheme,
		"http_proto":  req.Proto,
		"http_method": req.Method,
		"remote_addr": req.RemoteAddr,
		"user_agent":  req.UserAgent(),
		"uri":         fmt.Sprintf("%s://%s%s", scheme, req.Host, req.RequestURI),
	})
	errLog.Error(err)
}
