package product

import (
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

func (p *dbRepository) Create(data *entity.Product) error {
	return p.db.Create(data).Error
}

func (p *dbRepository) FindAll() ([]entity.Product, error) {
	var data []entity.Product
	err := p.db.Preload("Image").Find(&data).Error

	return data, err
}

func (p *dbRepository) Find(spec specifications.I) ([]entity.Product, error) {
	var data []entity.Product
	err := spec.GormQuery(p.db.Preload("Image")).Find(&data).Error

	return data, err
}

func (p *dbRepository) Get(spec specifications.I) (entity.Product, error) {
	var data entity.Product
	err := spec.GormQuery(p.db.Preload("Image")).First(&data).Error

	return data, err
}
