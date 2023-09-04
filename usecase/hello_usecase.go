package usecase

import (
	"context"
	"time"

	"github.com/swaingotnochill/tempmee/domain"
)

type helloUsecase struct {
	helloRepository domain.HelloRepository
	contextTimeout time.Duration
}

func NewHelloUseCase(helloRepository domain.HelloRepository, timeout time.Duration) domain.HelloUsecase {
	return &helloUsecase {
		helloRepository: helloRepository,
		contextTimeout: timeout,
	}
}

func (hu *helloUsecase) PrintHello(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, hu.contextTimeout)
	defer cancel()

	return hu.helloRepository.PrintHello(ctx)
}