package todosservices

import (
	todosmodels "e-routine/modules/todos/models"
	todosrepositories "e-routine/modules/todos/repositories"
	"e-routine/shared/app_error"

	"github.com/gofiber/fiber/v2"
)

func GetTodosService() ([]todosmodels.Todo, error) {
	return todosrepositories.GetTodos()
}

func GetTodosController(ctx *fiber.Ctx) error {
	todos, err := GetTodosService()
	if err != nil {
		return app_error.HandleHttp(ctx, err)
	}

	return ctx.JSON(todos)
}
