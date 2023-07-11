package http

import (
	"e-routine/shared/providers/env"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Start() {
	var app = fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		var variable = "Hello"
		var returnStatement = variable + "Irwin"

		return ctx.SendString(returnStatement)
	})
	log.Fatal(app.Listen(":" + env.Get(env.Port)))
}
