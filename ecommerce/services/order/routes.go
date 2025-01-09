package order

import (
	"ecommerce/types"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
)

type Handler struct {
	store types.OrderStore
}

func NewHandler(store types.OrderStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /orders", h.handleGetOrders)
	router.HandleFunc("PATCH /orders/{orderId}/status", h.handleStatusUpdate)
}

func (h *Handler) handleStatusUpdate(w http.ResponseWriter, r *http.Request) {
	var payload types.UpdateOrderStatusPayload

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

	orderID, _ := strconv.Atoi(payload.ID)

	err = h.store.UpdateOrder(orderID, payload.Status)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) handleGetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.store.GetOrdersByUserId(1)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string][]types.Order{
		"orders": orders,
	})
}
