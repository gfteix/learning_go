package types

import (
	"context"
	"time"
)

type CartItemPayload struct {
	ProductID int `json:"productId" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type AddressPayload struct {
	Street     string `json:"street" validate:"required"`
	City       string `json:"city" validate:"required"`
	State      string `json:"state" validate:"required"`
	PostalCode string `json:"postalCode" validate:"required"`
	Country    string `json:"country" validate:"required"`
}

type CartCheckoutPayload struct {
	Items     []CartItemPayload `json:"items" validate:"required"`
	AddressId *int              `json:"addressId,omitempty"` // Optional if using a saved address
	Address   *AddressPayload   `json:"address,omitempty"`   // Optional if providing a new address
}

type LoginUserPayload struct {
	Email    string `json:"email"  validate:"required,email"`
	Password string `json:"password"  validate:"required,min=3,max=130"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName"  validate:"required"`
	Email     string `json:"email"  validate:"required,email"`
	Password  string `json:"password"  validate:"required"`
}

type UpdateOrderStatusPayload struct {
	Status string `json:"status" validate:"required"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(user User) error
}

type ProductStore interface {
	GetProducts() ([]Product, error)
	GetProductsByIDs(ids []int) ([]Product, error)
	UpdateProduct(Product) error
}

type AddressStore interface {
	GetAddress(addressId int) (*Address, error)
	GetAddressesByUserID(userID int) ([]Address, error)
	CreateAddress(address Address, userID int) (int, error)
}

type OrderStore interface {
	CreateOrder(ctx context.Context, productMap map[int]Product, cart CartCheckoutPayload, userID int) (int, float64, error)
	CreateOrderItem(OrderItem) error
	GetOrdersByUserId(userID int) ([]Order, error)
	GetOrder(id int) (*Order, error)
	UpdateOrder(orderID int, status string) error
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Address struct {
	ID         int    `json:"id"`
	UserID     int    `json:"userId"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Order struct {
	ID         int         `json:"id"`
	UserID     int         `json:"userID"`
	Total      float64     `json:"total"`
	Status     string      `json:"status"`
	AddressId  int         `json:"addressId"`
	CreatedAt  time.Time   `json:"createdAt"`
	OrderItems []OrderItem `json:"orderItems"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"orderID"`
	ProductID int       `json:"productID"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}
