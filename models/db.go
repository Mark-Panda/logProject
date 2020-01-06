package models

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	"log"
	"net/url"
	"os"
)

type Config struct {
	DataSource string `toml:"dataSource"`
}

func GetDB() *gorm.DB {
	dataSource := dataSource()
	db, err := gorm.Open("mysql", dataSource) //root:@(127.0.0.1)/blog?charset=utf8&parseTime=True&loc=Local
	if err != nil {
		panic(err)
	}
	db.SingularTable(true) //表生成结尾不带s
	return db
}

func dataSource() string {
	address := os.Getenv("CLEARDB_DATABASE_URL")
	if address == "" {
		return defaultDataSource()
	}
	url, err := url.Parse(address)
	if err != nil {
		log.Fatal(err)
	}
	return getProductDataSource(url.User.String(), url.Host, url.Path)
}

func getProductDataSource(user, host, databasePath string) string {
	fmt.Println("commmmmm")

	var dataSourceParams = [...]string{
		user,
		"@tcp(",
		host,
		":3306",
		")",
		databasePath,
		//"?parseTime=true",
	}
	var dataSource string
	for _, dataSourceParam := range dataSourceParams {
		dataSource += dataSourceParam
	}
	fmt.Println("dataSource---", dataSource)
	return dataSource
}

func defaultDataSource() string {
	var config Config
	_, err := toml.DecodeFile("mysql/conf.toml", &config)
	if err != nil {
		panic(err)
	}
	return config.DataSource
}