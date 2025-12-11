package main

import (
	"3-validation-api/configs"
	"3-validation-api/internal/auth"
	"net/http"
)

func main() {
	config := configs.GetConfig()
	router := http.NewServeMux()
	auth.NewVerifyHandler(router, auth.AuthHandlerDeps{Config: config})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	server.ListenAndServe()
}
