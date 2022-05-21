package main

import (
	"github.com/gofiber/fiber/v2"
	"gofiber04-jwtauth/route"
)

func main() {
	app := fiber.New()

	//创建路由
	route.SetupRoute(app)

	app.Listen(":8080")
}
