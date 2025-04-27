package product

import (
	"context"

	"github.com/hanhnham91/order-service/entity"
	"github.com/hanhnham91/order-service/repository/specifications"
)

type Repository interface {
	FindAll(ctx context.Context) ([]entity.Product, error)
	Find(ctx context.Context, spec specifications.I) ([]entity.Product, error)
	Get(ctx context.Context, spec specifications.I) (entity.Product, error)
	Create(ctx context.Context, data *entity.Product) error
}
