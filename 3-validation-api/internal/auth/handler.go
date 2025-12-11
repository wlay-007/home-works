package auth

import (
	"3-validation-api/configs"
	"3-validation-api/pkg/email"
	"3-validation-api/pkg/file"
	"3-validation-api/pkg/hash"
	"3-validation-api/pkg/req"
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

type VerificationRecord struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func NewVerifyHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &VerifyHandler{
		deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[VerifyPayload](w, r)
		if err != nil {
			return
		}
		hashStr, err := hash.GenerateHash(6)
		if err != nil {
			return
		}
		fmt.Println(body.Email)
		fmt.Println(hashStr)
		email.SendEmail(hashStr, handler.Email, handler.Password, body.Email)

		rec := VerificationRecord{
			Email: body.Email,
			Hash:  hashStr,
		}

		err = file.SaveVerification(rec)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		user, err := file.LoadVerification[VerificationRecord]()
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		if hash != user.Hash {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, "Email verified !", http.StatusOK)
	}
}
