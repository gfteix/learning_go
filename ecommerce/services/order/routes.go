package order

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
	store     types.OrderStore
	userStore types.UserStore
}

func NewHandler(store types.OrderStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:     store,
		userStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /orders", auth.WithJWTAuth(h.handleGetOrders, h.userStore))
	router.HandleFunc("PATCH /orders/{orderId}/status", auth.WithJWTAuth(h.handleStatusUpdate, h.userStore))
}

func (h *Handler) handleStatusUpdate(w http.ResponseWriter, r *http.Request) {
	orderID, _ := strconv.Atoi(r.PathValue("orderId"))

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

	order, err := h.store.GetOrder(orderID)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if order == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid order id"))
		return
	}

	err = h.store.UpdateOrder(orderID, payload.Status)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) handleGetOrders(w http.ResponseWriter, r *http.Request) {
	context := r.Context()
	userID := auth.GetUserIDFromContext(context)

	orders, err := h.store.GetOrdersByUserId(userID)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string][]types.Order{
		"orders": orders,
	})
}
