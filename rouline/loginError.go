package rouline

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type LoginError struct {}

func (self LoginError) RegisterRoute(g *echo.Group)  {
	g.GET("/", self.ReadError)
	g.POST("/login",self.WriteError)
}

func (LoginError) ReadError(ctx echo.Context)  error{

}

func (LoginError) WriteError(ctx echo.Context) error {

}