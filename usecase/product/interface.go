package product

import (
	"context"

	"github.com/hanhnham91/order-service/entity"
)

type IProductFindUseCase interface {
	Execute(ctx context.Context) ([]entity.Product, error)
}

type IProductGetUseCase interface {
	Execute(ctx context.Context, id int64) (entity.Product, error)
}
