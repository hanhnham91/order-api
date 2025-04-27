package auth

import (
	"github.com/labstack/echo/v4"
)

type Route struct{}

func Init(group *echo.Group) {
	r := &Route{}
	g := group.Group("/auth")

	g.POST("/connect-social", r.connectsocial)
}
