package rouline

import (
	"fmt"
	"github.com/labstack/echo"
)

type Mechanism struct {}

func (mech Mechanism) MechanismRegisterRoute(group *echo.Group) {
	fmt.Println("进来就是是是")
	group.POST("/mechanism/produce", mech.Produce)
}

func (mech Mechanism) Produce(ctx echo.Context) error {
	fmt.Println("sddffff")
	fmt.Println("准备插入机构数据库", ctx.FormValue("table"))

	return nil
}
