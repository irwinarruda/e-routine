package http

import (
	"database/sql"
	"e-routine/shared/providers/env"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		instance, err := sql.Open("postgres", env.Get(env.DbURL))
		if err != nil {
			return ctx.SendString("Error opening database" + err.Error())
		}
		rows, err := instance.Query("SELECT * FROM todos")
		if err != nil || rows.Err() != nil {
			return ctx.SendString("Error querying todos")
		}
		defer func() {
			rows.Close()
		}()
		columns, err := rows.Columns()
		if err != nil {
			return ctx.SendString("Error getting columns" + err.Error())
		}
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for rows.Next() {
			for i := range columns {
				valuePtrs[i] = &values[i]
			}
			err := rows.Scan(valuePtrs...)
			if err != nil {
				return ctx.SendString("Error scanning row" + err.Error())
			}
		}
		if err := rows.Err(); err != nil {
			return ctx.SendString("Error iterating rows" + err.Error())
		}

		return ctx.SendString(fmt.Sprintf("%+v", values))
	})
	log.Fatal(app.Listen(":" + env.Get(env.Port)))
}
