package dtos

import "time"

type AuthenticateUserDto struct {
	Email    string
	Password string
}

type TokenDataDto struct {
	Id       int
	ExpireIn time.Time
	Audience string
}

type AuthenticatedUserDto struct {
	Id          int
	AccessToken string
	Kind        string
	ExpireIn    time.Time
}
