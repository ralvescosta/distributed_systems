package entities

import "time"

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
