package address

import (
	"ecommerce/services/auth"
	"ecommerce/types"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store types.AddressStore
}

func NewHandler(store types.AddressStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /users/{id}/addresses", h.handleCreateAddress)
	router.HandleFunc("GET /users/{id}/addresses", h.handleGetAddresses)
}

func (h *Handler) handleCreateAddress(w http.ResponseWriter, r *http.Request) {
	authUserId := auth.GetUserIDFromContext(r.Context())

	var payload types.AddressPayload

	err := utils.ParseJson(r, &payload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	id, err := h.store.CreateAddress(types.Address{
		Street:     payload.Street,
		City:       payload.City,
		State:      payload.State,
		Country:    payload.Country,
		PostalCode: payload.PostalCode,
	}, authUserId)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]int{
		"id": id,
	})
}

func (h *Handler) handleGetAddresses(w http.ResponseWriter, r *http.Request) {
	pathUserID := r.PathValue("id")

	if pathUserID == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid path parameter"))
		return
	}

	authUserId := auth.GetUserIDFromContext(r.Context())
	userIDAsInt, _ := strconv.Atoi(pathUserID)

	if authUserId != userIDAsInt {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	addresses, err := h.store.GetAddressesByUserID(userIDAsInt)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string][]types.Address{
		"addresses": addresses,
	})
}
