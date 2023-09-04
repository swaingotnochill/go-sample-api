package repository

import (
	"context"

	"github.com/swaingotnochill/tempmee/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookRepository struct {
	database   *mongo.Database
	collection string
}

func NewBookRepository(db *mongo.Database, collection string) domain.BookRepository {
	return &bookRepository{
		database:   db,
		collection: collection,
	}
}

// DeleteBook implements domain.BookRepository.
func (br *bookRepository) DeleteBook(c context.Context, id string) error {
	collection := br.database.Collection(br.collection)
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return  err
	}

	filter := bson.M{"_id": idHex}
	_, err = collection.DeleteOne(c, filter)
	return err
}

// GetBooks implements domain.BookRepository.
func (br *bookRepository) GetBooks(c context.Context) ([]*domain.Book, error) {
	collection := br.database.Collection(br.collection)

	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var books []*domain.Book
	for cursor.Next(c) {
		var book domain.Book
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}

// GetBooksByID implements domain.BookRepository.
func (br *bookRepository) GetBooksByID(c context.Context, id string) (*domain.Book, error) {
	collection :=  br.database.Collection(br.collection)

	var book domain.Book

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &book, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&book)
	return &book, err
}

// SaveBook implements domain.BookRepository.
func (br *bookRepository) SaveBook(c context.Context, book *domain.Book) error {
	collection :=  br.database.Collection(br.collection)
	_, err := collection.InsertOne(c, book)
	return err
}

