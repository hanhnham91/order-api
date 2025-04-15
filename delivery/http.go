package http

import (
	"github.com/hanhnham91/order-service/delivery/order"
	"github.com/hanhnham91/order-service/delivery/product"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewHTTPHandler() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Secure())
	e.Use(middleware.RequestID())
	e.Use(middleware.Gzip())

	apiGroup := e.Group("/api")

	product.Init(apiGroup)
	order.Init(apiGroup)

	return e
}
