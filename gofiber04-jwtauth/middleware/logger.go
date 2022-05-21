package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"time"
)

//初始化fiber中间件
func InitLoggerMiddleware(app *fiber.App) {
	//日志
	app.Use(logger.New(logger.Config{
		//格式化日志的内容
		Format:     "[${time}] ${status} - ${ip}  ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05-0700",
		TimeZone:   "Asia/Shanghai",
		// 时间更新间隔
		// Optional. Default: 500 * time.Millisecond
		TimeInterval: 500 * time.Millisecond,
		//Output:       logFile, //定义了这个值，就会写入到文件中
		//Output: os.Stderr,
	}))
}
