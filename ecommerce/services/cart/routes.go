package cart

import (
	"ecommerce/services/auth"
	"ecommerce/types"
	"ecommerce/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
	userStore    types.UserStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:        store,
		productStore: productStore,
		userStore:    userStore,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore))
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())

	var cart types.CardCheckoutPayload

	if err := utils.ParseJson(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	productIDs, err := getCartItemsIDs(cart.Items)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	products, err := h.productStore.GetProductsByIDs(productIDs)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(products, cart.Items, userID)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"total":    totalPrice,
		"order_id": orderID,
	})
}
