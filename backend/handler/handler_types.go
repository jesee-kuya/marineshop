package handler

import (
	"github.com/jesee-kuya/marineshop/middleware"
	"github.com/jesee-kuya/marineshop/service"
)

type Authentication struct {
	AuthService service.AuthService
}

type Marineshop struct {
	Auth        Authentication
	AuthService service.AuthService
	Middleware  middleware.Middleware
}
