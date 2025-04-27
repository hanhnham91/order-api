package middleware

import (
	"github.com/hanhnham91/order-service/codetype"
	"github.com/hanhnham91/order-service/config"
	"github.com/hanhnham91/order-service/entity"
	"github.com/hanhnham91/order-service/repository/specifications"
	"github.com/hanhnham91/order-service/repository/user"
	"github.com/hanhnham91/order-service/util"
	pkgerror "github.com/hanhnham91/pkg/error"
	"github.com/labstack/echo/v4"
	"gitlab.com/pak-server/pkg/pakerror"
)

func Auth(userRepo user.Repository) func(next echo.HandlerFunc) echo.HandlerFunc {
	handlerFunc := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get(echo.HeaderAuthorization)
			jwt := util.NewJwt()

			claims, err := jwt.DecodeToken(token, config.GetConfig().TokenSecretKey)
			if err != nil {
				return util.Response.Error(c, pkgerror.ErrUnauthorized())
			}

			myUser, err := userRepo.Get(c.Request().Context(), specifications.UserByEmail(claims.Email, false))
			if err != nil {
				return util.Response.Error(c, pakerror.ErrInternalServerError(err))
			}

			if myUser.Status == codetype.StatusInactive || !myUser.IsVerified {
				return util.Response.Error(c, pkgerror.ErrUnauthorized())
			}

			c.Set(entity.MyUserClaim, &myUser)

			return next(c)
		}
	}

	return handlerFunc
}
