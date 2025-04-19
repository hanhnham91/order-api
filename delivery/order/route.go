package order

import (
	"github.com/hanhnham91/order-service/usecase"
	"github.com/hanhnham91/order-service/usecase/order"
	"github.com/labstack/echo/v4"
)

type Route struct {
	orderCreateUseCase order.IOrderCreateUseCase
}

func NewOrderRoute(orderCreateUC order.IOrderCreateUseCase) *Route {
	return &Route{
		orderCreateUseCase: orderCreateUC,
	}
}

func Init(group *echo.Group) {
	orderCreateUseCase := usecase.InjectOrderCreateUseCase()

	r := NewOrderRoute(orderCreateUseCase)
	g := group.Group("/order")

	g.POST("", r.create)
}
