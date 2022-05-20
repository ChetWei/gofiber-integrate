package models

import (
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model //继承 会默认填充 `created_at` 当前时间,`updated_at` 当前时间,`deleted_at`为空  字段

	Id   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	//注意 对于声明了默认值的字段，像 0、''、false 等零值是不会保存到数据库。
	//您需要使用指针类型或 Scanner/Valuer 来避免这个问题
	//
	Age      *int   `gorm:"default:18"`   //使用默认值，并且使用指针，就可以将默认值插入
	Active   bool   `gorm:"default:true"` //bool 会自动转为 0 1
	Password string `json:"password"`
}

//Hook 是在创建、查询、更新、删除等操作之前、之后调用的函数。

/*钩子函数,当操作 User对应的表，会执行*/

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("before create")
	return
}
