package dtos

import "time"

type SignInDto struct {
	Email    string
	Password string
}

type TokenDataDto struct {
	Id       int
	ExpireIn time.Time
	Audience string
}

type SessionDto struct {
	Id          int
	AccessToken string
	Kind        string
	ExpireIn    time.Time
}
