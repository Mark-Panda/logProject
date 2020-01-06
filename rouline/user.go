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
	group.POST("/login",user.WriteError)
}


func (user User) ReadError(ctx echo.Context)  error{
	//var userInfo []models.User
	phone := ctx.QueryParam("phone")
	use := models.User{}
	fmt.Println("phone", phone)
	use.FindOneByOps(phone)
	//fmt.Println("-------",userInfo)
	return ctx.JSON(http.StatusOK, "sss")
}

func (user User) WriteError(ctx echo.Context) error {
	return nil
}