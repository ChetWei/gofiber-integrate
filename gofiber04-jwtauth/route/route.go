package route

import (
	"github.com/gofiber/fiber/v2"
	"gofiber04-jwtauth/handler"
	"gofiber04-jwtauth/middleware"
)

func SetupRoute(app *fiber.App) {

	//日志中间件
	middleware.InitLoggerMiddleware(app)

	app.Post("/login", handler.Login)

	//jwt认证中间件,以下的路由都要经过中间件认证
	middleware.InitJwtMiddleware(app)

	app.Get("/hello", handler.Hello)
	app.Get("/hi", handler.Hi)
}
