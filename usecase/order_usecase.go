package usecase

import (
	"context"
	"time"

	"github.com/swaingotnochill/tempmee/domain"
)

type orderUsecase struct {
	orderRepository domain.OrderRepository
	contextTimeout  time.Duration
}

func NewOrderUseCase(orderRepository domain.OrderRepository, timeout time.Duration) domain.OrderUseCase {
	return &orderUsecase{
		orderRepository: orderRepository,
		contextTimeout:  timeout,
	}
}

func (ou *orderUsecase) GetAllOrders(c context.Context) ([]*domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.GetAllOrders(ctx)

}

// CreateOrder implements domain.OrderUseCase.
func (ou *orderUsecase) CreateOrder(c context.Context, order *domain.Order) error {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.CreateOrder(ctx, order)
}

// GetOrderByID implements domain.OrderUseCase.
func (ou *orderUsecase) GetOrderByID(c context.Context, id string) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.GetOrderByID(ctx, id)
}

// GetOrdersByCustomerID implements domain.OrderUseCase.
func (ou *orderUsecase) GetOrdersByCustomerID(c context.Context, id string) ([]*domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepository.GetOrdersByCustomerID(ctx, id)
}
