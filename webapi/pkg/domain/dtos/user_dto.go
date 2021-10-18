package dtos

import "time"

type CreateUserDto struct {
	Name     string
	Email    string
	Password string
}

type CreatedUserDto struct {
	Id          int
	Name        string
	Email       string
	CreatedAt   time.Time
	AccessToken string
	Kind        string
	ExpireIn    time.Time
}
