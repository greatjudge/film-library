package main

import (
	"filmlibr/internal/handlers"
	"filmlibr/internal/logger"
	"filmlibr/internal/middleware"
	service "filmlibr/internal/service/actor"
	"filmlibr/internal/session"

	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func DummyHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic("log init error")
	}
	defer func() {
		err = zapLogger.Sync()
		if err != nil {
			fmt.Println(err)
		}
	}()

	logr := logger.NewLogger(zapLogger.Sugar())

	actorService := service.NewActorService()
	actorHandler := handlers.NewActorHandler(actorService)

	sessionManger := session.NewSessionManagerImpl()

	// Handlers TODO:
	mux := http.NewServeMux()

	mux.HandleFunc("/actor/{id}", actorHandler.Actor) // Get, Put, Patch, Delete
	mux.HandleFunc("/actors", actorHandler.List)      // Get, Post

	mux.HandleFunc("/film/{id}/{actor_id}", DummyHandler) // Post, Delete
	mux.HandleFunc("/film/{id}", DummyHandler)            // Get, Patch, Put, Delete
	mux.HandleFunc("/films", DummyHandler)                // Get, Post

	// query -> Panic -> AccessLog -> Auth -> AccessControl -> JSONContentCheck
	appHandler := middleware.Panic(logr, mux)
	appHandler = middleware.AccessLog(logr, appHandler)
	appHandler = middleware.Auth(sessionManger, appHandler)
	appHandler = middleware.AccessControl(appHandler)
	appHandler = middleware.JSONContentCheck(appHandler)

	// TODO: use env vars for host:port
	err = http.ListenAndServe(":8000", appHandler)
	if err != nil {
		panic(err)
	}
}
