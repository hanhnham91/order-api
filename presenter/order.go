package presenter

import (
	"github.com/hanhnham91/order-service/entity"
)

//nolint:tagliatelle
type OrderResponse struct {
	ID         int64               `json:"id"`
	CouponCode string              `json:"couponCode"`
	Amount     float64             `json:"amount"`
	OrderItems []OrderItemResponse `json:"orderItems"`
	Products   []ProductResponse   `json:"products"`
}

//nolint:tagliatelle
type OrderItemResponse struct {
	ProductID int64 `json:"productId"`
	Quantity  int   `json:"quantity"`
}

func FormOrdeeResponse(myOrder *entity.Order, products []entity.Product) OrderResponse {
	resp := OrderResponse{
		ID:         myOrder.ID,
		CouponCode: myOrder.CouponCode,
		Amount:     myOrder.Amount,
	}

	resp.OrderItems = FormOrderItemsResponse(myOrder.OrderItems)

	resp.Products = FormProductsResponse(products)

	return resp
}

func FormOrderItemsResponse(orderItems []entity.OrderItem) []OrderItemResponse {
	resp := make([]OrderItemResponse, len(orderItems))

	for i := range orderItems {
		resp[i] = OrderItemResponse{
			ProductID: orderItems[i].ProductID,
			Quantity:  orderItems[i].Quantity,
		}
	}

	return resp
}
