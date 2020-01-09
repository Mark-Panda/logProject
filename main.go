package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"loggerProject/models"
	"loggerProject/rouline"
	"net/http"
	"os"
	"time"
)




func main()  {
	e := echo.New()
	e.Use(middleware.Logger())  //每个http请求记录
	fileAppend, err := os.OpenFile("./systemHttpLog.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		fmt.Println("打开文件失败", err)
		os.Exit(1)
	}
	//将每个http请求写到文件日志中，若关闭下面配置，则只打印到控制台上
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: fileAppend,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	//e.Use(middleware.CSRF())
	g := e.Group("/v1")
	rouline.RegisterRoutes(g)

	models.InitClient()  //启动Redis
	models.InitTables() //同步表结构
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
