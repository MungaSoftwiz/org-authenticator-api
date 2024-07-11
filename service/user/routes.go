package user

import (
	"fmt"
	"log"
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
	router.HandleFunc("/auth/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/auth/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/api/users/{userId}", h.handleGetUser).Methods("GET")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload

	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"status":     "Bad request",
			"message":    "Invalid payload",
			"statusCode": 400,
		})
		return
	}

	// Validate the payload
	if err := utils.Validate.Struct(&payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{
			"status":     "Bad request",
			"message":    "Validation failed",
			"statusCode": 400,
		})
		return
	}

	user, err := h.storage.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteJSON(w, http.StatusNotFound, map[string]interface{}{
			"status":     "Not found",
			"message":    "User not found",
			"statusCode": http.StatusNotFound,
		})
		return
	}

	log.Printf("User found: %+v\n", user)

	// check if password match
	if !auth.CheckPasswordHash(payload.Password, user.Password) {
		utils.WriteJSON(w, http.StatusUnauthorized, map[string]interface{}{
			"status":     "Unauthorized",
			"message":    "Invalid credentials",
			"statusCode": http.StatusUnauthorized,
		})
		return
	}

	token, err := auth.GenerateToken(strconv.Itoa(user.ID))
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":     "error",
			"message":    "Could not generate token",
			"statusCode": 500,
		})
		return
	}

	log.Printf("Login successful for user with email %s\n", payload.Email)

	response := map[string]interface{}{
		"status":  "success",
		"message": "Login successful",
		"data": map[string]interface{}{
			"accessToken": token,
			"user": map[string]interface{}{
				"userId":    user.ID,
				"firstName": user.FirstName,
				"lastName":  user.LastName,
				"email":     user.Email,
				"phone":     user.Phone,
			},
		},
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// handle register
	var payload types.RegisterUserPayload
	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, struct {
			Status     string `json:"status"`
			Message    string `json:"message"`
			StatusCode int    `json:"statusCode"`
		}{
			Status:     "Bad request",
			Message:    "Registration unsuccessful",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(&payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// check if user already exists
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

	// Create a new user object
	newUser := types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
		Phone:     payload.Phone,
	}

	// save the new user to the database
	err = h.storage.CreateUser(newUser)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	token, err := auth.GenerateToken(strconv.Itoa(newUser.ID))
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]interface{}{
			"status":     "error",
			"message":    "Could not generate token",
			"statusCode": 500,
		})
		return
	}

	// Successful response
	response := struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  "success",
		Message: "Registration successful",
		Data: struct {
			AccessToken string     `json:"accessToken"`
			User        types.User `json:"user"`
		}{
			AccessToken: token,
			User:        newUser,
		},
	}

	utils.WriteJSON(w, http.StatusCreated, response)
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	str, ok := params["userId"]
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
		utils.WriteJSON(w, http.StatusNotFound, struct {
			Status     string `json:"status"`
			Message    string `json:"message"`
			StatusCode int    `json:"statusCode"`
		}{
			Status:     "Not found",
			Message:    "User not found",
			StatusCode: http.StatusNotFound,
		})
		return
	}

	// Prepare successful response
	response := struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  "success",
		Message: "User details retrieved successfully",
		Data: struct {
			UserID    string `json:"userId"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Email     string `json:"email"`
			Phone     string `json:"phone"`
		}{
			UserID:    strconv.Itoa(user.ID), // Convert user.ID to string
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
		},
	}

	utils.WriteJSON(w, http.StatusOK, response)
}
