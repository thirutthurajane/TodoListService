package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"todoservice/configs"
	"todoservice/logger"
	"todoservice/route"
)

func main() {

	mylog := logrus.New()
	mylog.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RealIP)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(logger.NewStructuredLogger(mylog))
	route.Setup(router)
	port := ""
	switch configs.Configuration.Server.Env {
	case "Develop":
		port = configs.Configuration.Server.Develop.Port
	case "Production":
		port = configs.Configuration.Server.Production.Port
	default:
		port = ":9999"
	}

	mylog.Println("Server Started On Port " + port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		mylog.Error(errors.Cause(err).Error())
		return
	}
}
