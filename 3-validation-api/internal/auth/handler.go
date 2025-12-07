package auth

import (
	"3-validation-api/configs"
	"3-validation-api/pkg/res"
	"fmt"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}
type VerifyHandler struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &VerifyHandler{
		deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("POST /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Send")
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Verify")
	}
}
