package product

import (
	"context"

	"github.com/hanhnham91/order-service/entity"
	"github.com/hanhnham91/order-service/repository/specifications"
	"gorm.io/gorm"
)

func NewRepo(getDB func() *gorm.DB) Repository {
	return &dbRepository{getDB()}
}

type dbRepository struct {
	db *gorm.DB
}

func (p *dbRepository) Create(ctx context.Context, data *entity.Product) error {
	return p.db.WithContext(ctx).Create(data).Error
}

func (p *dbRepository) FindAll(ctx context.Context) ([]entity.Product, error) {
	var data []entity.Product
	err := p.db.WithContext(ctx).Preload("Image").Find(&data).Error

	return data, err
}

func (p *dbRepository) Find(ctx context.Context, spec specifications.I) ([]entity.Product, error) {
	var data []entity.Product
	err := spec.GormQuery(p.db.WithContext(ctx).Preload("Image")).Find(&data).Error

	return data, err
}

func (p *dbRepository) Get(ctx context.Context, spec specifications.I) (entity.Product, error) {
	var data entity.Product
	err := spec.GormQuery(p.db.WithContext(ctx).Preload("Image")).First(&data).Error

	return data, err
}
