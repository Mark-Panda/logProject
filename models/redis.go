package models

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的redisdb变量
var Redisdb *redis.Client

// 初始化连接
func InitClient() (err error) {
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = Redisdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}


func redisExample() {
	err := Redisdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := Redisdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := Redisdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}