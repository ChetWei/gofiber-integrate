package models

import (
	"time"
)

//存放对应数据库表的实体模型
//以及一些数据库相关的操作

//Book model
type Book struct {
	Id        uint      `gorm:"primary_key"`
	Guid      string    `gorm:"not null"` // 设置字段为非空并唯一
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time //gorm约定使用了CreatedAt，UpdatedAt字段会自动填充当前时间
	UpdatedAt time.Time
	DeletedAt *time.Time //使用指针，可以传递一个空值进去
}
