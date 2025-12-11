package auth

type VerifyPayload struct {
	Email string `json:"email" validate:"required,email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
