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

type ResponsInfo struct {
	Code int64 `json:"code"`
	Msg string `json:"msg"`
}

func RegisterRoutes(g *echo.Group) {
	//API插入日志统一入口
	g.POST("/unifiedStdin", func(ctx echo.Context) error {
		cs := &RequestType{}
		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			fmt.Println("读取HTTPbody失败", err)
			return err
		}
		json.Unmarshal(body, &cs)
		level := cs.Level
		switch level {
			case "formal":
				//fmt.Println("正常业务", ctx)
				InsertAnylog(cs)
				resInfo := &ResponsInfo{
					Code: 200,
					Msg:  "success",
				}
				return ctx.JSON(http.StatusOK, resInfo)
			case "catch":
				//fmt.Println("异常业务")
				AbnormalLog(cs)
				resInfo := &ResponsInfo{
					Code: 200,
					Msg:  "success",
				}
				return ctx.JSON(http.StatusOK, resInfo)
			default:
				//fmt.Println("默认正常业务")
				InsertAnylog(cs)
				resInfo := &ResponsInfo{
					Code: 200,
					Msg:  "success",
				}
				return ctx.JSON(http.StatusOK, resInfo)
		}
	})

	/*
	业务路由
	 */
	new(User).UserRegisterRoute(g)
	new(Mechanism).MechanismRegisterRoute(g)
}

//正常日志
func InsertAnylog(r *RequestType)  {
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
func httpLocal(r *RequestType) error {
	url := "http://127.0.0.1:1323/v1/" + r.Table + "/produce"
	infoJson, _ := json.Marshal(r)
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

	resp,err :=http.DefaultClient.Do(req)
	if err!=nil{
		fmt.Printf("post数据请求error:%v\n",err)
		return err
	}else {
		respBody,_ :=ioutil.ReadAll(resp.Body)
		fmt.Printf("post数据请求successful:%v\n",string(respBody))
		return nil
	}

}