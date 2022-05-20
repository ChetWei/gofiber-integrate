package routes

import (
	"github.com/gofiber/fiber/v2"
	"gofiber02-gorm/routes/api/v1"
)

//初始化路由
func InitRouter(app *fiber.App) {
	//book
	app.Get("/hello", api.Hello)
	app.Get("/book", api.GetBook)

	//user
	app.Post("/user", api.AddUser)
	app.Post("/users", api.AddUsers)
}
