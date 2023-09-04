package repository

import (
	"context"
	"fmt"

	"github.com/swaingotnochill/tempmee/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type helloRepository struct {
	database *mongo.Database
	collection string
}

func NewHelloRepository(db *mongo.Database, collection string) domain.HelloRepository {
	return &helloRepository{
		database: db,
		collection: collection,
	}
}

func (hr *helloRepository) PrintHello(c context.Context) error {
	fmt.Println("Hello")
	return nil
}