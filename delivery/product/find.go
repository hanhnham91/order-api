package product

import (
	"github.com/hanhnham91/order-service/presenter"
	"github.com/hanhnham91/order-service/usecase"
	"github.com/hanhnham91/order-service/util"
	"github.com/labstack/echo/v4"
)

func (r *Route) find(c echo.Context) error {
	uc := usecase.InjectProductFindUseCase()

	products, err := uc.Execute(c.Request().Context())
	if err != nil {
		return util.Response.Error(c, err)
	}

	return util.Response.Success(c, presenter.FormProductsResponse(products))
}
