package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/swaingotnochill/tempmee/domain"
	"github.com/swaingotnochill/tempmee/domain/mocks"
	"github.com/swaingotnochill/tempmee/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetBooks(t *testing.T) {
	mockBookRepository := new(mocks.BookRepository)
	
	t.Run("success", func (t *testing.T) {

		mockBook := &domain.Book {
			ID : primitive.NewObjectID(),
			Title : "Dummy Title",
			Author : "Dummy Author",
			Price : 1.23,
		}

		mockBookList := make([]*domain.Book, 0)
		mockBookList = append(mockBookList, mockBook)


		mockBookRepository.On("GetBooks", mock.Anything).Return(mockBookList, nil).Once()

		bu := usecase.NewBookUsecase(mockBookRepository, time.Second * 2)

		list , err := bu.GetBooks(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.Len(t, list, len(mockBookList))

		mockBookRepository.AssertExpectations(t)
	})

	t.Run("error", func (t *testing.T) {
		mockBookRepository.On("GetBooks", mock.Anything).Return(nil, errors.New("Unexpected")).Once()

		bu := usecase.NewBookUsecase(mockBookRepository, time.Second * 2)

		list, err := bu.GetBooks(context.Background())

		assert.Error(t, err)
		assert.Nil(t, list)

		mockBookRepository.AssertExpectations(t)
	})
}


func TestGetBooksByID(t *testing.T) {
	mockBookRepository := new(mocks.BookRepository)

	bookObjectID := primitive.NewObjectID()
	bookID := bookObjectID.Hex()


	t.Run("success", func (t *testing.T) {
		mockBook := &domain.Book{
			ID : bookObjectID,
			Title: "Dummy Title",
			Author : "Dummy Author",
			Price: 1.23,
		}

		mockBookRepository.On("GetBooksByID", mock.Anything, bookID).Return(mockBook, nil).Once()

		bu := usecase.NewBookUsecase(mockBookRepository, time.Second * 2)

		book, err := bu.GetBooksByID(context.Background(), bookID)

		assert.NoError(t, err)
		assert.NotNil(t, book)
		assert.Equal(t, book, mockBook)

		mockBookRepository.AssertExpectations(t)
	})

	t.Run("error", func (t *testing.T) {
		mockBookRepository.On("GetBooksByID", mock.Anything, bookID).Return(nil, errors.New("Unexpected")).Once()

		bu := usecase.NewBookUsecase(mockBookRepository, time.Second * 2)

		book, err := bu.GetBooksByID(context.Background(), bookID)

		assert.Error(t , err)
		assert.Nil(t, book)

		mockBookRepository.AssertExpectations(t)
	})
}

func TestDeleteBook(t *testing.T) {
	mockBookRepository := new(mocks.BookRepository)

	bookObjectID := primitive.NewObjectID()
	bookID := bookObjectID.Hex()

	t.Run("success", func (t *testing.T) {
		mockBookRepository.On("DeleteBook", mock.Anything, bookID).Return(nil).Once()

		bu := usecase.NewBookUsecase(mockBookRepository, time.Second * 2)

		err := bu.DeleteBook(context.Background(), bookID)

		assert.NoError(t, err)

		mockBookRepository.AssertExpectations(t)
	})
	t.Run("error", func (t *testing.T) {
		mockBookRepository.On("DeleteBook", mock.Anything, bookID).Return(errors.New("Unexpected")).Once()

		bu := usecase.NewBookUsecase(mockBookRepository, time.Second * 2)

		err := bu.DeleteBook(context.Background(), bookID)

		assert.Error(t, err)

		mockBookRepository.AssertExpectations(t)
	})
}

func TestSaveBook(t *testing.T) {
	mockBookRepository := new(mocks.BookRepository)

	bookObjectID := primitive.NewObjectID()

	mockBook := &domain.Book {
		ID : bookObjectID,
		Title : "Dummy Title",
		Author: "Dummy Author",
		Price : 1.23,
	}

	t.Run("success", func (t *testing.T){
		mockBookRepository.On("SaveBook", mock.Anything, mockBook).Return(nil).Once()

		bu := usecase.NewBookUsecase(mockBookRepository, time.Second * 2)

		err := bu.SaveBook(context.Background(), mockBook)
		assert.NoError(t, err)

		mockBookRepository.AssertExpectations(t)
	})
	t.Run("error", func (t *testing.T){
		mockBookRepository.On("SaveBook", mock.Anything, mockBook).Return(errors.New("Unexpected")).Once()

		bu := usecase.NewBookUsecase(mockBookRepository, time.Second * 2)

		err := bu.SaveBook(context.Background(), mockBook)

		assert.Error(t, err)

		mockBookRepository.AssertExpectations(t)
	})
}