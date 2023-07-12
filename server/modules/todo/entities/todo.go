package entities

import "time"

type Todo struct {
	id          int64
	title       string
	description string
	weekDays    []int64
	createdAt   time.Time
	updatedAt   time.Time
}

func (this Todo) ToEntitie(rowValues []interface{}) {
	if id, ok := rowValues[0].(int64); ok {
		this.id = id
	} else {
		panic("Error converting id to int64")
	}

	if title, ok := rowValues[1].(string); ok {
		this.title = title
	} else {
		panic("Error converting title to string")
	}

	if description, ok := rowValues[2].(string); ok {
		this.description = description
	} else {
		panic("Error converting description to string")
	}

	// if weekDays, ok := rowValues[3].([]int64); ok {
	// 	this.weekDays = weekDays
	// } else {
	// 	panic("Error converting weekDays to []int64")
	// }

	// if createdAt, ok := rowValues[4].(time.Time); ok {
	// 	this.createdAt = createdAt
	// } else {
	// 	panic("Error converting createdAt to time.Time")
	// }

	// if updatedAt, ok := rowValues[5].(time.Time); ok {
	// 	this.updatedAt = updatedAt
	// } else {
	// 	panic("Error converting updatedAt to time.Time")
	// }
}
