package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBook = "books"
)

type Book struct {
	ID       primitive.ObjectID  `bson:"_id"`
	Title    string  `bson:"title"`
	Author   string  `bson:"author"`
	Price    float64 `bson:"price"`
}

type BookRepository interface {
	GetBooks(c context.Context) ([]*Book, error)
	GetBooksByID(c context.Context, id string) (*Book, error)
	SaveBook(c context.Context, book *Book) error
	DeleteBook(c context.Context, id string) error
}

type BookUseCase interface {
	GetBooks(c context.Context) ([]*Book, error)
	GetBooksByID(c context.Context, id string) (*Book, error)
	SaveBook(c context.Context, book *Book) error
	DeleteBook(c context.Context, id string) error
}
