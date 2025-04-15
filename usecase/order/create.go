package order

import (
	"context"

	"github.com/hanhnham91/order-service/entity"
	"github.com/hanhnham91/order-service/payload"
	"github.com/hanhnham91/order-service/repository/order"
	"github.com/hanhnham91/order-service/repository/product"
	"github.com/hanhnham91/order-service/repository/specifications"
	pkgerror "github.com/hanhnham91/pkg/error"
)

type orderCreateUseCase struct {
	productRepo product.Repository
	orderRepo   order.Repository
}

func NewOrderCreateUseCase(productRepo product.Repository, orderRepo order.Repository) IOrderCreateUseCase {
	return &orderCreateUseCase{
		productRepo: productRepo,
		orderRepo:   orderRepo,
	}
}

func (u *orderCreateUseCase) Execute(
	_ context.Context,
	req payload.CreateOrderRequest,
) (*entity.Order, []entity.Product, error) {
	var (
		myOrder = &entity.Order{
			CouponCode: req.CouponCode,
		}
		productIDs = make([]int64, 0)
		mProducts  = make(map[int64]entity.Product)
		total      float64
	)

	for i := range req.Items {
		productIDs = append(productIDs, req.Items[i].ProductID)
	}

	products, err := u.productRepo.Find(specifications.ProductByIDs(productIDs))
	if err != nil {
		return nil, nil, pkgerror.ErrInternalServerError(err)
	}

	for _, product := range products {
		mProducts[product.ID] = product
	}

	for i := range req.Items {
		myProduct, ok := mProducts[req.Items[i].ProductID]
		if !ok {
			return nil, nil, pkgerror.ErrNotFound("Product not found")
		}

		total += float64(req.Items[i].Quantity) * myProduct.Price

		myOrder.OrderItems = append(myOrder.OrderItems, entity.OrderItem{
			ProductID: myProduct.ID,
			Quantity:  req.Items[i].Quantity,
		})
	}

	myOrder.Amount = total

	if err = u.orderRepo.Create(myOrder); err != nil {
		return nil, nil, pkgerror.ErrInternalServerError(err)
	}

	return myOrder, products, nil
}
