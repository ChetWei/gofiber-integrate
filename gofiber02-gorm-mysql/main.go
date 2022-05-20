package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"

	"gofiber02-gorm/conf"
	"gofiber02-gorm/models"
	"gofiber02-gorm/routes"
)

func getConnetInfo() string {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DatabaseSetting.User, conf.DatabaseSetting.Password, conf.DatabaseSetting.Host, conf.DatabaseSetting.Name)

	return dsn
}

//初始化
func init() {
	//1.初始化配置文件
	conf.Setup()
	//2.初始化数据库连接实例
	models.Setup()
}

func main() {

	//1.创建fiber实例
	app := fiber.New()

	//2.配置fiber中间件
	routes.InitMiddleware(app)

	//3.初始化路由 包含中间件的配置
	routes.InitRouter(app)

	//4.开启监听
	log.Fatal(app.Listen(":8080"))
}
