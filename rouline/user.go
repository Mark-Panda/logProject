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
	//err := models.Redisdb.Set("phone", "100", 0).Err()
	//if err != nil {
	//	fmt.Printf("set score failed, err:%v\n", err)
	//	return nil
	//}
	userInfo := use.FindOneByOps(phone)
	return ctx.JSON(http.StatusOK, userInfo)
}

func (user User) WriteError(ctx echo.Context) error {
	phone := ctx.QueryParam("phone")
	name := ctx.QueryParam("name")
	pwd := ctx.QueryParam("pwd")
	use := models.User{}
	userInfo := use.CreateUser(phone, name, pwd)
	return ctx.JSON(http.StatusOK, userInfo)
}