package dtos

import "time"

type AuthenticationDto struct {
	Email    string
	Password string
}

type UserAuthenticated struct {
	AccessToken string
	Kind        string
	ExpiredAt   time.Time
}
