package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MungaSoftwiz/org-authenticator-api/service/auth"
	"github.com/MungaSoftwiz/org-authenticator-api/types"
	"github.com/MungaSoftwiz/org-authenticator-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	storage types.UserStorage
}

func NewHandler(storage types.UserStorage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// handle login
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// handle register
	var payload types.RegisterUserPayload
	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(&payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	_, err := h.storage.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteJSON(w, http.StatusBadRequest, "user already exists")
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	user := types.User{
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Password:  hashedPassword,
		Phone:     payload.Phone,
	}

	err = h.storage.CreateUser(user)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "user created successfully")
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["userID"]
	if !ok {
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("missing user ID"))
		return
	}

	userID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("invalid user ID"))
		return
	}

	user, err := h.storage.GetUserByID(userID)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}
