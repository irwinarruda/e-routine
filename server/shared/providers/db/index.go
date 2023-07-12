package db

import (
	"database/sql"
	"e-routine/shared/app_error"
	"e-routine/shared/providers/env"
	"fmt"
)

func Query(instance *sql.DB, query string, args ...any) (*sql.Rows, error) {
	rows, err := instance.Query(query, args...)

	return rows, err
}

func ExecuteSelect[T DBEntitie](rows *sql.Rows) ([]T, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, app_error.New("Error getting columns "+err.Error(), app_error.InternalServerError)
	}
	finalValues := []T{}
	rowValues := make([]interface{}, len(columns))
	rowValuePtrs := make([]interface{}, len(columns))
	fmt.Println("Before for loop")
	for rows.Next() {
		fmt.Println("Inside for loop")
		for i := range columns {
			rowValuePtrs[i] = &rowValues[i]
		}
		fmt.Println("Before scan")
		err := rows.Scan(rowValuePtrs...)
		if err != nil {
			return nil, app_error.New("Error scanning row "+err.Error(), app_error.InternalServerError)
		}
		fmt.Println("After scan")
		newValue := new(T)
		(*newValue).ToEntitie(rowValues)
		finalValues = append(finalValues, *newValue)
	}
	fmt.Println("After for loop")
	fmt.Println(finalValues)

	if err := rows.Err(); err != nil {
		return nil, app_error.New("Error iterating rows "+err.Error(), app_error.InternalServerError)
	}

	return finalValues, nil
}

func Start() (*sql.DB, error) {
	instance, err := sql.Open("postgres", env.Get(env.DbURL))
	if err != nil {
		return nil, app_error.New("Error opening database"+err.Error(), app_error.NoStatus)
	}

	return instance, nil
}
