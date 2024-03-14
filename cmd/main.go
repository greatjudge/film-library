package main

import (
	"filmlibr/internal/logger"
	"filmlibr/internal/middleware"
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

	sessionManger := session.NewSessionManagerImpl()

	// Handlers
	mux := http.NewServeMux()

	mux.HandleFunc("/actors", DummyHandler)     // Get, Post
	mux.HandleFunc("/actor/{id}", DummyHandler) // Get, Put, Patch, Delete

	mux.HandleFunc("/films", DummyHandler)                // Get, Post
	mux.HandleFunc("/film/{id}", DummyHandler)            // Get, Patch, Put, Delete
	mux.HandleFunc("/film/{id}/{actor_id}", DummyHandler) // Post, Delete

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
