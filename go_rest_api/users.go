package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var errEmailRequired = errors.New("email is required")
var errFirstNameRequired = errors.New("first name is required")
var errLastNameRequired = errors.New("last name is required")
var errPasswordRequired = errors.New("password is required")

type UserService struct {
	store Store
}

func NewUserService(s Store) *UserService {
	return &UserService{store: s}
}

func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", s.handleUserRegister).Methods("POST")
	r.HandleFunc("/users/login", s.handleUserLogin).Methods("POST")
}

func (s *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var payload *User
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if err := validateUserPayload(payload); err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	hashedPassword, err := HashPassword(payload.Password)
	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
		return
	}
	payload.Password = hashedPassword

	u, err := s.store.CreateUser(payload)
	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
		return
	}

	token, err := createAndSetAuthCookie(u.Id, w)
	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating session"})
		return
	}

	WriteJson(w, http.StatusCreated, token)
}

func (s *UserService) handleUserLogin(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var payload *LoginPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if err := validateLoginPayload(payload); err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
		return
	}

	user, err := s.store.GetUserByEmail(payload.Email)

	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusNotFound, ErrorResponse{Error: "User Not Found"})
		return
	}

	isValid := ComparePasswords(user.Password, payload.Password)

	if !isValid {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid Login"})
		return
	}

	token, err := createAndSetAuthCookie(user.Id, w)
	if err != nil {
		log.Println(err.Error())
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating session"})
		return
	}

	WriteJson(w, http.StatusOK, token)
}

func validateLoginPayload(payload *LoginPayload) error {
	if payload.Email == "" {
		return errEmailRequired
	}

	if payload.Password == "" {
		return errPasswordRequired
	}

	return nil
}

func validateUserPayload(user *User) error {
	if user.Email == "" {
		return errEmailRequired
	}

	if user.FirstName == "" {
		return errFirstNameRequired
	}

	if user.LastName == "" {
		return errLastNameRequired
	}

	if user.Password == "" {
		return errPasswordRequired
	}

	return nil
}

func createAndSetAuthCookie(id int64, w http.ResponseWriter) (string, error) {
	secret := []byte(Envs.JWTSecret)
	token, err := CreateJWT(secret, id)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	return token, nil
}
