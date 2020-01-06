package models

import "fmt"

type User struct {
	ID   string // 名为`ID`的字段会默认作为表的主键
	Name string
	Phone string
}

func (use *User) FindOneByOps(phone string) *User {
	fmt.Println("参数", phone)
	user := User{}
	if err := GetDB().Where("phone = ?", phone); err != nil {
		return nil
	}
	fmt.Println("数据库", user)
	return &user
}



// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
//db.SingularTable(true)