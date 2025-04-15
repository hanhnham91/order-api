package usecase

import (
	"github.com/hanhnham91/order-service/client/sql"
	orderRepo "github.com/hanhnham91/order-service/repository/order"
	productRepo "github.com/hanhnham91/order-service/repository/product"
	"github.com/hanhnham91/order-service/usecase/order"
	"github.com/hanhnham91/order-service/usecase/product"
)

func InjectProductFindUseCase() product.IProductFindUseCase {
	return product.NewProductFindUseCase(
		productRepo.NewRepo(sql.GetClient),
	)
}

func InjectProductGetUseCase() product.IProductGetUseCase {
	return product.NewProductGetByIDUseCase(
		productRepo.NewRepo(sql.GetClient),
	)
}

func InjectOrderCreateUseCase() order.IOrderCreateUseCase {
	return order.NewOrderCreateUseCase(
		productRepo.NewRepo(sql.GetClient),
		orderRepo.NewRepo(sql.GetClient),
	)
}
