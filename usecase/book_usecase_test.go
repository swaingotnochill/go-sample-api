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