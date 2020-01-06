package rouline

import (
	"fmt"
	"github.com/labstack/echo"
	"loggerProject/models"
	"net/http"
)

type User struct {}

func (user User) UserRegisterRoute(group *echo.Group) {
	group.GET("/use", user.ReadError)
	group.GET("/add", user.WriteError)
	group.POST("/login",user.WriteError)
}


func (user User) ReadError(ctx echo.Context)  error{
	//var userInfo []models.User
	phone := ctx.QueryParam("phone")
	use := models.User{}
	fmt.Println("phone", phone)
	userInfo := use.FindOneByOps(phone)
	return ctx.JSON(http.StatusOK, userInfo)
}

func (user User) WriteError(ctx echo.Context) error {
	phone := ctx.QueryParam("phone")
	name := ctx.QueryParam("name")
	use := models.User{}
	userInfo := use.CreateUser(phone, name)
	return ctx.JSON(http.StatusOK, userInfo)
}