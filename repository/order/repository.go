package order

import (
	"context"

	"github.com/hanhnham91/order-service/entity"
)

type Repository interface {
	Create(ctx context.Context, data *entity.Order) error
}
