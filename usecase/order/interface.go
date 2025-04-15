package order

import (
	"context"

	"github.com/hanhnham91/order-service/entity"
	"github.com/hanhnham91/order-service/payload"
)

type IOrderCreateUseCase interface {
	Execute(ctx context.Context, req payload.CreateOrderRequest) (*entity.Order, []entity.Product, error)
}
