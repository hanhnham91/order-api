package order

import (
	"github.com/labstack/echo/v4"
)

type Route struct{}

func Init(group *echo.Group) {
	r := &Route{}
	g := group.Group("/order")

	g.POST("", r.create)
}
