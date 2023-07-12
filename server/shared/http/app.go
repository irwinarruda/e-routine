package http

import (
	"e-routine/modules/todo/entities"
	"e-routine/shared/providers/db"
	"e-routine/shared/providers/env"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		instance, err := db.Start()
		if err != nil {
			return ctx.SendString("Error opening database" + err.Error())
		}
		rows, err := db.Query(instance, "SELECT * FROM todos")
		if err != nil || rows.Err() != nil {
			return ctx.SendString("Error querying todos")
		}
		defer rows.Close()

		todos, err := db.ExecuteSelect[entities.Todo](rows)
		if err != nil {
			return ctx.SendString("Error getting todos " + err.Error())
		}

		return ctx.SendString(fmt.Sprintf("%+v", todos))
	})
	log.Fatal(app.Listen(":" + env.Get(env.Port)))
}
