package todosusecases

import (
	todosmodels "e-routine/modules/todos/models"
	"e-routine/shared/app_error"
	"e-routine/shared/providers/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func open() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(env.Get(env.DbURL)))
	if err != nil {
		return nil, app_error.NewPrivate("Failed to connect database", app_error.InternalServerError)
	}

	return db, nil
}

func GetTodos() ([]todosmodels.Todo, error) {
	db, err := open()
	if err != nil {
		return nil, err
	}

	todos := []todosmodels.Todo{}
	db.Table("todos").Find(&todos)

	return todos, nil
}
