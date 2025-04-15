package product

import (
	"strconv"

	"github.com/hanhnham91/order-service/presenter"
	"github.com/hanhnham91/order-service/usecase"
	"github.com/hanhnham91/order-service/util"
	pkgerror "github.com/hanhnham91/pkg/error"
	"github.com/labstack/echo/v4"
)

func (r *Route) get(c echo.Context) error {
	uc := usecase.InjectProductGetUseCase()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return util.Response.Error(c, pkgerror.ErrBadRequest("Invalid product_id"))
	}

	myProduct, err := uc.Execute(c.Request().Context(), id)
	if err != nil {
		return util.Response.Error(c, err)
	}

	resp := presenter.FormProductResponse(myProduct)

	return util.Response.Success(c, resp)
}
