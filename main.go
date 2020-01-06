package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"loggerProject/models"
	"loggerProject/rouline"
	"net/http"
	"time"
)




func main()  {
	//initDB()
	e := echo.New()
	e.Use(middleware.Logger())  //每个http请求记录
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.CSRF())
	g := e.Group("/v1")
	rouline.RegisterRoutes(g)

	models.InitClient()
	/**
	自定义启动方式
	 */
	s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))

}

//func initDB()  {
//	db := models.GetDB()
//	db.AutoMigrate(models.Login{})
//}
