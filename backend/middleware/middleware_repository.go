package middleware

import "github.com/gin-gonic/gin"

type Middleware interface {
	RouteChecker() gin.HandlerFunc
	AuthMiddleware() gin.HandlerFunc
	RequireRole(roles ...string) gin.HandlerFunc
}
