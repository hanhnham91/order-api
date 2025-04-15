package product

import (
	"github.com/labstack/echo/v4"
)

type Route struct{}

func Init(group *echo.Group) {
	r := &Route{}
	g := group.Group("/product")

	g.GET("", r.find)
	g.GET("/:id", r.get)
}
