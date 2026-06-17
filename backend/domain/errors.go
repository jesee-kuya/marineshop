package domain

import "errors"

var (
	ErrEmailInUse         = errors.New("email already in use")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserNotFound       = errors.New("user not found")
)
