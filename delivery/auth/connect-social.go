package auth

import (
	"github.com/hanhnham91/order-service/payload"
	"github.com/hanhnham91/order-service/usecase"
	"github.com/hanhnham91/order-service/util"
	pkgerror "github.com/hanhnham91/pkg/error"
	"github.com/labstack/echo/v4"
)

func (r *Route) connectsocial(c echo.Context) error {
	var req payload.AuthFirebaseRequest

	if err := c.Bind(&req); err != nil {
		return util.Response.Error(c, pkgerror.ErrBadRequest("Error binding request"))
	}

	if err := req.Validate(); err != nil {
		return util.Response.Error(c, pkgerror.ErrValidation(err))
	}

	uc := usecase.InjectAuthFirebaseUseCase()

	token, err := uc.Execute(c.Request().Context(), req)
	if err != nil {
		return util.Response.Error(c, err)
	}

	return util.Response.Success(c, map[string]string{"token": token})
}
