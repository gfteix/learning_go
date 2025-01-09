package types

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
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

type CartItem struct {
	ProductID int
	Quantity  int
}

type CardCheckoutPayload struct {
	Items []CartItem `json:"items" validate:"required"`
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

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"` // not a good practice to store the quantity in the product table
	CreatedAt   time.Time `json:"createdAt"`
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userID"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"orderID"`
	ProductID int       `json:"productID"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}
