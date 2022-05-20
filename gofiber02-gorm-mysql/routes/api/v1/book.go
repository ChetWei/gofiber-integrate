package api

import (
	"github.com/gofiber/fiber/v2"
	"gofiber02-gorm/models"
)

/*
配置路由的处理方法
*/

func Hello(ctx *fiber.Ctx) error {
	return ctx.SendString("hello world!")
}

func GetBook(ctx *fiber.Ctx) error {
	result := &models.Book{}
	err := models.DBConn.First(result).Error
	if err != nil {
		return err
	}
	return ctx.JSON(result)
}
