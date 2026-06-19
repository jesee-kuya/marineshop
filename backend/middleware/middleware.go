package middleware

import "github.com/jesee-kuya/marineshop/domain"

func NewMiddleware(jwt *domain.JWTConfig) Middleware {
	return &MiddlewareStruct{JWT: jwt}
}
