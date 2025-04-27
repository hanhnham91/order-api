package user

import (
	"context"

	"github.com/hanhnham91/order-service/entity"
	"github.com/hanhnham91/order-service/repository/specifications"
)

type Repository interface {
	Get(ctx context.Context, spec specifications.I) (entity.User, error)
	Create(ctx context.Context, data *entity.User) error
	Update(ctx context.Context, data *entity.User) error
	CreateOrDoUpdate(ctx context.Context, data *entity.User, updatedFields ...string) error
}
