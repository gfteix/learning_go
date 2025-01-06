package user

import (
	"ecommerce/services/auth"
	"ecommerce/types"
	"ecommerce/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /login", h.handleLogin)
	router.HandleFunc("POST /register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	log.Print("handleLogin")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	log.Print("handleRegister")

	var payload types.RegisterUserPayload

	err := utils.ParseJson(r, &payload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}
	// check if the user exists

	u, err := h.store.GetUserByEmail(payload.Email)

	if u != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	hashedPassword, err := auth.HashPassowrd(payload.Password)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)

}
