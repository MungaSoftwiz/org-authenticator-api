package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/auth/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/auth/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
		// handle login
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
		// handle register
}