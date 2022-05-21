package handler

import (
	"github.com/gofiber/fiber/v2"
	"gofiber04-jwtauth/middleware"
)

func Login(ctx *fiber.Ctx) error {
	user := ctx.FormValue("user")
	password := ctx.FormValue("password")
	// Throws Unauthorized error
	if user != "john" || password != "doe" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	//创建token
	token, err := middleware.CreateToken(middleware.BaseClaims{user})
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.JSON(fiber.Map{"token": token})
}

func Hello(ctx *fiber.Ctx) error {
	c := "hello func"
	return ctx.SendString(c)
}

func Hi(ctx *fiber.Ctx) error {
	c := "hi func"
	return ctx.SendString(c)
}
