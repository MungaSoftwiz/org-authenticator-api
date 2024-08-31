package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MungaSoftwiz/org-authenticator-api/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStorage := &mockUserStorage{}
	handler := NewHandler(userStorage)

	t.Run("fails if the user ID is not a number", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/user/abc", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/user/{userID}", handler.handleGetUser).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("handles get user by ID", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/user/42", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/user/{userID}", handler.handleGetUser).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("fails if user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "john",
			LastName:  "doe",
			Email:     "asd",
			Password:  "invalid",
			Phone:     "123-245",
		}
		marshalledPayload, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalledPayload))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/request", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expects status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("user register successful", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "john",
			LastName:  "doe",
			Email:     "asd",
			Password:  "example@email.com",
			Phone:     "123-245",
		}
		marshalledPayload, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalledPayload))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/request", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expects status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStorage struct{}

func (m *mockUserStorage) UpdateUser(u types.User) error {
	return nil
}

func (m *mockUserStorage) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, fmt.Errorf("user not found")
}

func (m *mockUserStorage) CreateUser(u types.User) error {
	return nil
}

func (m *mockUserStorage) GetUserByID(id int) (*types.User, error) {
	return &types.User{}, nil
}
