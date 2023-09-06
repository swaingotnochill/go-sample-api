package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionOrders = "order"
)

type OrderRequest struct {
	Books []string `form:"books" binding:"required"`
}

type Order struct {
	ID          primitive.ObjectID   `bson:"_id"`
	CustomerID  primitive.ObjectID   `bson:"customer_id"`
	BookIDs     []string `bson:"book_ids"`
	TotalAmount float64              `bson:"total_amount"`
	CreatedAt   time.Time            `bson:"created_at"`
}

type OrderRepository interface {
	GetAllOrders(c context.Context) ([]*Order, error)
	CreateOrder(c context.Context, order *Order) error
	GetOrderByID(c context.Context, id string) (*Order, error)
	GetOrdersByCustomerID(c context.Context, id string) ([]*Order, error)
	GetBookById(c context.Context, id string) (*Book, error)
}

type OrderUseCase interface {
	GetAllOrders(c context.Context) ([]*Order, error)
	CreateOrder(c context.Context, order *Order) error
	GetOrderByID(c context.Context, id string) (*Order, error)
	GetOrdersByCustomerID(c context.Context, id string) ([]*Order, error)
	GetBookById(c context.Context, id string) (*Book, error)
}
