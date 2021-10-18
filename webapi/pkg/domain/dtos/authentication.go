package dtos

import "time"

type AuthenticationDto struct {
	Email    string
	Password string
}

type TokenDataDto struct {
	Id       uint32
	ExpireIn time.Time
	Audience string
}

type AuthenticatedUserDto struct {
	AccessToken string
	Kind        string
	ExpireIn    time.Time
}
