package http

import (
	todosusecases "e-routine/modules/todos/usecases"
	"e-routine/shared/providers/env"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()
	app.Get("/todos", todosusecases.GetTodosController)
	log.Fatal(app.Listen(":" + env.Get(env.Port)))
}
