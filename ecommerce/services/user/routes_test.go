package user

import (
	"bytes"
	"ecommerce/types"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockUserStore struct{}

func TestUserServiceHandlers(t *testing.T) {

	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "LastName",
			Email:     "invalid",
			Password:  "asw",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}

	})

	t.Run("should correctly register the user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "LastName",
			Email:     "valid@email.com",
			Password:  "asw",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		log.Print(rr.Body.String())
		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil // returning no user
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}
