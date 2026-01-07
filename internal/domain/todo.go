package domain

import "time"

type Todo struct {
	ID        int64
	Title     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
