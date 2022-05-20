package api

import (
	"github.com/gofiber/fiber/v2"
	. "gofiber02-gorm/models"
	"log"
)

//单行插入
func AddUser(ctx *fiber.Ctx) error {
	user := new(User)
	err := ctx.BodyParser(user)
	log.Println(user)
	if err == nil {
		//通过传递过来的数据插入数据库
		//result := DBConn.Create(&user)

		//指定字段插入
		//result := DBConn.Select("Name", "Password").Create(&user)

		//插入时忽略指定的字段
		result := DBConn.Omit("Id").Create(&user)

		//跳过钩子方法执行
		//DBConn.Session(&gorm.Session{SkipHooks: true}).Create(&user)

		//根据Map创建
		//result := DBConn.Model(&User{}).Create(map[string]interface{}{
		//	"Name": "jinzhu", "password": "helloworld",
		//})

		return ctx.JSON(result.RowsAffected)
	}
	return err
}

//批量插入
func AddUsers(ctx *fiber.Ctx) error {
	// >>> 1.方法一 一批次插入
	//将一个 slice 传递给 Create 方法。 GORM 将生成单独一条SQL语句来插入所有数据，并回填主键的值
	//var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	//INSERT INTO `user` (`name`,`password`) VALUES ('jinzhu1',''),('jinzhu2',''),('jinzhu3','')
	//result := DBConn.Create(&users)

	// >>> 2.方法二 分批创建时，你可以指定每批的数量
	/*	var users2 = []User{{Name: "jinzhu_1"}, {Name: "jinzhu_2"}, {Name: "jinzhu_3"}, {Name: "jinzhu_4"}, {Name: "jinzhu_5"}, {Name: "jinzhu_6"}, {Name: "jinzhu_7"}}
		// 每批插入5行，自动分批数
		result := DBConn.CreateInBatches(users2, 5)*/

	// >>> 3. 根据map批量创建，也会整合成一条sql语句插入
	result := DBConn.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Active": true},
		{"Name": "jinzhu_2", "Active": false},
	})
	//  根据 map 创建记录时，association 不会被调用，且主键也不会自动填充

	return ctx.JSON(result.RowsAffected)
}
