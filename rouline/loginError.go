package rouline

import (
	"fmt"
	"github.com/labstack/echo"
	"loggerProject/models"
	"loggerProject/mysql"
	"net/http"
)

type LoginError struct {}

func (self LoginError) RegisterRoute(g *echo.Group)  {
	g.GET("/", self.ReadError)
	g.POST("/login",self.WriteError)
}

func (LoginError) ReadError(ctx echo.Context)  error{
	var logInfo []models.Login
	phone := ctx.QueryParam("phone")
	fmt.Println("phone", phone)
	db := mysql.GetDB()
	defer db.Close()
	err := db.Where("phone = ?", phone).Find(logInfo)
	if err != nil {
		fmt.Println("errpr", err)
		return nil
	}
	fmt.Println("-------",logInfo)
	return ctx.JSON(http.StatusOK, "sss")
}

func (LoginError) WriteError(ctx echo.Context) error {
	return nil
}