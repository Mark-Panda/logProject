package rouline

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"loggerProject/models"
	"net/http"
	"time"
)

type User struct {}

func (user User) UserRegisterRoute(group *echo.Group) {
	group.GET("/use", user.ReadError)
	group.GET("/add", user.WriteError)
	group.POST("/getToken", user.GetJwtToken)
	group.GET("/login", user.Login)
}

func (user User) Login(ctx echo.Context) error  {
	fmt.Println("come come")
	//fmt.Println("ccccc", ctx.Get("user").(*jwt.Token) )
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
	name := ctx.FormValue("name")
	pwd := ctx.FormValue("pwd")
	fmt.Println("入参", name, pwd)
	use := models.User{}
	sign := use.CheckUserInfo(name, pwd)
	if sign {
		token := jwt.New(jwt.SigningMethodHS256)
		//nameKey := name + "secrty"
		//fmt.Println("hhhh", nameKey)
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["admin"] = true
		//claims["exp"] = time.Now().Add(time.Minute * 3).Unix()
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}
