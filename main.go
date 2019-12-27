package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"loggerProject/models"
	"loggerProject/mysql"
	"loggerProject/rouline"
)

func main()  {
	initDB()
	e := echo.New()
	e.Use(middleware.Logger())  //每个http请求记录
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.CSRF())
	g := e.Group("/v1")
	rouline.RegisterRoutes(g)

}

func initDB()  {
	db := mysql.GetDB()
	db.AutoMigrate(models.Login{})
}
