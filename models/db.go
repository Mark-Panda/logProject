package models

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/url"
	"os"
)

type Config struct {
	DataSource string `toml:"dataSource"`
}

func GetDB() *gorm.DB {
	dataSource := dataSource()
	db, err := gorm.Open("mysql", dataSource)
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

/*
同步表结构 gorm可以支持自动迁移，也就是自动的表结构迁移，只会创建表，补充缺少的列，缺少的索引。但并不会更改已经存在的列类型，也不会删除不再用的列，这样设计的目的是为了保护已存在的数据。可以同时针对多个表进行迁移设置。
 */
func InitTables()  {
	GetDB().AutoMigrate(&User{}, &Login{}, &Mechanism{})

}