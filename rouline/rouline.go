package rouline

import "github.com/labstack/echo"

func RegisterRoutes(g *echo.Group) {
	new(LoginError).RegisterRoute(g)
}