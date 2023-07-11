package entities

import "time"

type Todo struct {
	id          int
	title       string
	description string
	weekDays    []int
	createdAt   time.Time
	updatedAt   time.Time
}
