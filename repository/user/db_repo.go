package user

import (
	"context"

	"github.com/hanhnham91/order-service/entity"
	"github.com/hanhnham91/order-service/repository/specifications"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewRepo(getDB func() *gorm.DB) Repository {
	return &dbRepository{getDB()}
}

type dbRepository struct {
	db *gorm.DB
}

func (p *dbRepository) Create(ctx context.Context, data *entity.User) error {
	return p.db.WithContext(ctx).Create(data).Error
}

func (p *dbRepository) Get(ctx context.Context, spec specifications.I) (entity.User, error) {
	var data entity.User
	err := spec.GormQuery(p.db.WithContext(ctx)).First(&data).Error

	return data, err
}

func (p *dbRepository) Update(ctx context.Context, data *entity.User) error {
	return p.db.WithContext(ctx).Save(data).Error
}

func (p *dbRepository) CreateOrDoUpdate(ctx context.Context, data *entity.User, updatedFields ...string) error {
	return p.db.
		WithContext(ctx).
		Clauses(clause.OnConflict{
			DoNothing: len(updatedFields) == 0,
			DoUpdates: clause.AssignmentColumns(updatedFields),
		}).
		Create(&data).
		Error
}
