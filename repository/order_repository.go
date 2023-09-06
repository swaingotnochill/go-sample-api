package repository

import (
	"context"

	"github.com/swaingotnochill/tempmee/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderRepository struct {
	database   *mongo.Database
	collection string
}

func NewOrderRepository(db *mongo.Database, collection string) domain.OrderRepository {
	return &orderRepository{
		database:   db,
		collection: collection,
	}
}

func (or *orderRepository) GetAllOrders(c context.Context) ([]*domain.Order, error) {
	collection := or.database.Collection(or.collection)

	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var orders []*domain.Order
	for cursor.Next(c) {
		var order domain.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}

// CreateOrder implements domain.OrderRepository.
func (or *orderRepository) CreateOrder(c context.Context, order *domain.Order) error {
	collection := or.database.Collection(or.collection)

	_, err := collection.InsertOne(c, order)
	return err
}

// GetOrderByID implements domain.OrderRepository.
func (or *orderRepository) GetOrderByID(c context.Context, id string) (*domain.Order, error) {
	collection := or.database.Collection(or.collection)

	var order domain.Order

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &order, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&order)
	return &order, err
}

// GetOrdersByCustomerID implements domain.OrderRepository.
func (or *orderRepository) GetOrdersByCustomerID(c context.Context, id string) ([]*domain.Order, error) {
	collection := or.database.Collection(or.collection)

	var orders []*domain.Order
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return orders, err
	}

	cursor, err := collection.Find(c, bson.M{"customer_id": idHex})
	if err != nil {
		return orders, err
	}
	defer cursor.Close(c)

	for cursor.Next(c) {
		var order domain.Order
		if err := cursor.Decode(&order); err != nil {
			return orders, err
		}
		orders = append(orders, &order)
	}
	return orders, nil
}
