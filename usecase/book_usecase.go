package usecase

import (
	"context"
	"time"

	"github.com/swaingotnochill/tempmee/domain"
)

type bookUsecase struct {
	bookRepository domain.BookRepository
	contextTimeout time.Duration
}

func NewBookUsecase(bookRepository domain.BookRepository, timeout time.Duration) domain.BookUseCase {
	return &bookUsecase{
		bookRepository: bookRepository,
		contextTimeout: timeout,
	}
}

// DeleteBook implements domain.BookUseCase.
func (bu *bookUsecase) DeleteBook(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.bookRepository.DeleteBook(ctx, id)
}

// GetBooks implements domain.BookUseCase.
func (bu *bookUsecase) GetBooks(c context.Context) ([]*domain.Book, error) {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.bookRepository.GetBooks(ctx)
}

// GetBooksByID implements domain.BookUseCase.
func (bu *bookUsecase) GetBooksByID(c context.Context, id string) (*domain.Book, error) {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.bookRepository.GetBooksByID(ctx, id)
}

// SaveBook implements domain.BookUseCase.
func (bu *bookUsecase) SaveBook(c context.Context, book *domain.Book) error {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.bookRepository.SaveBook(ctx, book)
}
