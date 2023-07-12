package todosmodels

import "time"

type Todo struct {
	ID          int64
	Title       string
	Description string
	WeekDays    []int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
