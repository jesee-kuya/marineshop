package middleware

type Middleware interface {
	RouteChecker() gin.HandlerFunc
}