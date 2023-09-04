package domain

import "context"

type HelloUsecase interface {
	PrintHello(c context.Context) error
}

type HelloRepository interface {
	PrintHello(c context.Context) error
}
