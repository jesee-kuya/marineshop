package middleware

import "github.com/jesee-kuya/marineshop/domain"

const ClaimsKey = "claims"

type MiddlewareStruct struct {
	JWT *domain.JWTConfig
}
