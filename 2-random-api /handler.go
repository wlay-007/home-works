package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

type RandomHandler struct{}

func NewRandomHandler(router *http.ServeMux) {
	handler := &RandomHandler{}
	router.HandleFunc("/random", handler.Random())
}
func (h *RandomHandler) Random() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		n := rand.Intn(6) + 1
		s := strconv.Itoa(n)
		w.Write([]byte(s))
	}
}
