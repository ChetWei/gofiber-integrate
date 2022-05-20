package main

import (
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

func main() {

	app := fiber.New()

	// default
	//app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world!")
	})
	//
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://127.0.0.1/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	app.Listen(":8080")
}
