package rouline

import (
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
)



func RegisterRoutes(g *echo.Group) {
	//统一入口
	g.POST("/o1", func(ctx echo.Context) error {
		fmt.Println("上下文", ctx )
		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			fmt.Println("读取HTTPbody失败", err)
			return err
		}

		fmt.Println("json", string(body))
		level := ctx.FormValue("level")

		switch level {
			case "formal":
				fmt.Println("正常业务")
				InsertAnylog(ctx)
			case "catch":
				fmt.Println("异常业务")
			default:
				fmt.Println("正常业务")

		}
		return nil
	})
	//new(LoginError).RegisterRoute(g)
	new(User).UserRegisterRoute(g)
}

func InsertAnylog(c echo.Context)  {
	fmt.Println("comeoem")
	table := c.FormValue("table") //表业务
	fmt.Println("上下文", table)
}