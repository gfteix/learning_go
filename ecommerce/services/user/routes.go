package user

import (
	"ecommerce/types"
	"ecommerce/utils"
	"log"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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

	err := utils.ParseJson(r, payload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

}
