package order

import (
	"github.com/hanhnham91/order-service/payload"
	"github.com/hanhnham91/order-service/presenter"
	"github.com/hanhnham91/order-service/usecase"
	"github.com/hanhnham91/order-service/util"
	pkgerror "github.com/hanhnham91/pkg/error"
	"github.com/labstack/echo/v4"
)

func (r *Route) create(c echo.Context) error {
	var req payload.CreateOrderRequest

	if err := c.Bind(&req); err != nil {
		return util.Response.Error(c, pkgerror.ErrBadRequest("Error binding request"))
	}

	if err := req.Validate(); err != nil {
		return util.Response.Error(c, pkgerror.ErrValidation(err))
	}

	uc := usecase.InjectOrderCreateUseCase()

	myOrder, products, err := uc.Execute(c.Request().Context(), req)
	if err != nil {
		return util.Response.Error(c, err)
	}

	return util.Response.Success(c, presenter.FormOrdeeResponse(myOrder, products))
}
