package order

import (
	"context"

	"github.com/hanhnham91/order-service/entity"
	"gorm.io/gorm"
)

func NewRepo(getDB func() *gorm.DB) Repository {
	return &dbRepository{getDB()}
}

type dbRepository struct {
	db *gorm.DB
}

func (p *dbRepository) Create(ctx context.Context, data *entity.Order) error {
	return p.db.WithContext(ctx).Create(data).Error
}
