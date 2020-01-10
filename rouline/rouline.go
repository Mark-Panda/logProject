package rouline

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"strings"
)

//统一log入口，参数也要统一
type RequestType struct {
	Date string         //服务日志产生时间
	Msg string          //日志内容
	Name string
	MechanismId int64  //机构ID
	Level string       //日志等级
	Table string       //业务表
}


func RegisterRoutes(g *echo.Group) {
	//API插入日志统一入口
	g.POST("/o1", func(ctx echo.Context) error {
		cs := &RequestType{}
		fmt.Println("上下文", ctx.Request().Body )
		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			fmt.Println("读取HTTPbody失败", err)
			return err
		}
		fmt.Println("token是多少", ctx.Request().Header)
		tokenInfo := ctx.Request().Header
		info, ok := tokenInfo["Authorization"]
		if ok {
			fmt.Println("token的值为", info)
		}else {
			fmt.Println("没有token")
		}
		fmt.Println("json", string(body))
		json.Unmarshal(body, &cs)
		level := cs.Level
		switch level {
			case "formal":
				//fmt.Println("正常业务", ctx)
				InsertAnylog(cs)
			case "catch":
				//fmt.Println("异常业务")
				AbnormalLog(cs)
			default:
				//fmt.Println("默认正常业务")
				InsertAnylog(cs)
		}
		return nil
	})

	/*
	业务路由
	 */
	new(User).UserRegisterRoute(g)
	new(Mechanism).MechanismRegisterRoute(g)
}

//正常日志
func InsertAnylog(r *RequestType)  {
	//fmt.Println("正常日志level", r.Level)
	//fmt.Println("正常日志table", r.Table)
	//tableType := r.Table
	//g.POST("/" + tableType + "/produce", )
	httpLocal(r)
}

//异常日志
func AbnormalLog(r *RequestType)  {
	fmt.Println("异常日志level", r.Level)
	fmt.Println("异常日志table", r.Table)
}

/*
在统一日志写入接口处转换调用本地的一个HTTP请求，换到不同的路由上
 */
func httpLocal(r *RequestType)  {
	url := "http://127.0.0.1:1323/v1/" + r.Table + "/produce"
	infoJson, _ := json.Marshal(r)
	//fmt.Println("二进制吗", infoJson)
	stringJson := string(infoJson)
	fmt.Println("字符串吗", stringJson)
	req, _ := http.NewRequest("POST",url, strings.NewReader(stringJson))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "127.0.0.1:1323")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Length", "84")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("Authorization", token)

	resp,err :=http.DefaultClient.Do(req)
	if err!=nil{
		fmt.Printf("post数据请求error:%v\n",err)
	}else {
		fmt.Println("post数据请求successful.")
		respBody,_ :=ioutil.ReadAll(resp.Body)
		fmt.Printf("response data:%v\n",string(respBody))
	}

}