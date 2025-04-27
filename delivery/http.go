package http

import (
	"github.com/hanhnham91/order-service/client/sql"
	"github.com/hanhnham91/order-service/config"
	"github.com/hanhnham91/order-service/delivery/auth"
	"github.com/hanhnham91/order-service/delivery/order"
	"github.com/hanhnham91/order-service/delivery/product"
	md "github.com/hanhnham91/order-service/middleware"
	userRepo "github.com/hanhnham91/order-service/repository/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewHTTPHandler() *echo.Echo {
	cfg := config.GetConfig()
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Secure())
	e.Use(middleware.RequestID())
	e.Use(middleware.Gzip())

	if cfg.Stage == config.StageTypeLocal {
		e.Use(middleware.CORS())
	}

	apiGroup := e.Group("/api")

	auth.Init(apiGroup) // NoAuth

	authGroup := apiGroup.Group("", md.Auth(userRepo.NewRepo(sql.GetClient)))

	product.Init(authGroup)
	order.Init(authGroup)

	return e
}
