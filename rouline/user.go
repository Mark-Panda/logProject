package rouline

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"loggerProject/models"
	"net/http"
	"time"
)

type User struct {}

func (user User) UserRegisterRoute(group *echo.Group) {
	group.GET("/user/use", user.ReadError)
	group.GET("/user/add", user.WriteError)
	group.POST("/user/getToken", user.GetJwtToken)
	group.POST("/user/produce", user.Produce)

	group.Use(middleware.JWT([]byte("secret")))  //在验证token是需要在路由前设置和jwt一样的密钥
	group.GET("/user/login", user.Login)

}

func (user User) Produce(ctx echo.Context) error {
	fmt.Println("准备插入数据库", ctx.FormValue("name"))
	return nil
}

func (user User) Login(ctx echo.Context) error  {
	fmt.Println("come come")
	fmt.Println("ccccc", ctx.Get("user") )
	tokenInfo := ctx.Get("user").(*jwt.Token)
	claims := tokenInfo.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return ctx.String(http.StatusOK, "Welcome "+name+"!")
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

func (user User) GetJwtToken(ctx echo.Context) error  {
	fmt.Println("上下文", ctx)
	name := ctx.FormValue("name")
	pwd := ctx.FormValue("pwd")
	fmt.Println("入参", name, pwd)
	use := models.User{}
	sign := use.CheckUserInfo(name, pwd)
	if sign {
		token := jwt.New(jwt.SigningMethodHS256)
		nameKey := name + "secrty"
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = nameKey
		claims["admin"] = true
		//claims["exp"] = time.Now().Add(time.Minute * 3).Unix()
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()  //超时时间

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))  //密钥
		if err != nil {
			return err
		}
		ctx.Set("user", t)
		return ctx.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}


