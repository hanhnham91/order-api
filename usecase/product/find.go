package product

import (
	"context"

	"github.com/hanhnham91/order-service/entity"
	"github.com/hanhnham91/order-service/repository/product"
	pkgerror "github.com/hanhnham91/pkg/error"
)

type productFindUseCase struct {
	productRepo product.Repository
}

func NewProductFindUseCase(productRepo product.Repository) IProductFindUseCase {
	return &productFindUseCase{
		productRepo: productRepo,
	}
}

func (u *productFindUseCase) Execute(_ context.Context) ([]entity.Product, error) {
	products, err := u.productRepo.FindAll()
	if err != nil {
		return nil, pkgerror.ErrInternalServerError(err)
	}

	return products, nil
}
