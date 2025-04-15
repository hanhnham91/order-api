package product

import (
	"context"
	"errors"

	"github.com/hanhnham91/order-service/entity"
	"github.com/hanhnham91/order-service/repository/product"
	"github.com/hanhnham91/order-service/repository/specifications"
	pkgerror "github.com/hanhnham91/pkg/error"
	"gorm.io/gorm"
)

type productGetByIDUseCase struct {
	productRepo product.Repository
}

func NewProductGetByIDUseCase(productRepo product.Repository) IProductGetUseCase {
	return &productGetByIDUseCase{
		productRepo: productRepo,
	}
}

func (u *productGetByIDUseCase) Execute(_ context.Context, id int64) (entity.Product, error) {
	myProduct, err := u.productRepo.Get(specifications.ProductByID(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Product{}, pkgerror.ErrNotFound()
		}

		return entity.Product{}, pkgerror.ErrInternalServerError(err)
	}

	return myProduct, nil
}
