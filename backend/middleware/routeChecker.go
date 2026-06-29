package middleware

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

var allowedRoutes = map[string][]string{
	"/health":                             {"GET"},
	"/api/v1/auth/signup":                 {"POST"},
	"/api/v1/auth/login":                  {"POST"},
	"/api/v1/auth/reset-password":         {"POST"},
	"/api/v1/auth/change-password":        {"POST"},
	"/api/v1/auth/logout":                 {"POST"},
	"/api/v1/products":                    {"GET"},
	"/api/v1/products/:id":                {"GET"},
	"/api/v1/buyer/recommendations":       {"GET"},
	"/api/v1/buyer/profile":               {"GET"},
	"/api/v1/buyer/cart":                  {"POST", "GET"},
	"/api/v1/buyer/cart/:id":              {"DELETE"},
	"/api/v1/buyer/checkout":              {"POST"},
	"/api/v1/buyer/transactions":          {"GET"},
	"/api/v1/buyer/orders":                {"GET"},
	"/api/v1/admin/kyc":                   {"GET"},
	"/api/v1/admin/kyc/:id/approve":       {"PUT"},
	"/api/v1/admin/kyc/:id/reject":        {"PUT"},
	"/api/v1/seller/kyc":                  {"POST"},
	"/api/v1/seller/shop":                 {"POST"},
	"/api/v1/seller/payment":              {"POST"},
	"/api/v1/seller/payments":             {"GET"},
	"/api/v1/seller/products":             {"POST", "GET"},
	"/api/v1/seller/products/:id":         {"PUT", "DELETE"},
	"/api/v1/seller/withdraw":             {"POST"},
	"/api/v1/seller/transactions":         {"GET"},
	"/api/v1/seller/orders":               {"GET"},
	"/api/v1/seller/orders/:id":           {"PUT"},
	"/api/v1/seller/analytics":            {"GET"},
	"/api/v1/seller/profile":              {"GET"},
}

func (middleware *MiddlewareStruct) RouteChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedMethods, ok := allowedRoutes[c.FullPath()]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "route not found",
			})
			c.Abort()
			fmt.Println("not found")
			return
		}

		methodAllowed := slices.Contains(allowedMethods, c.Request.Method)

		if !methodAllowed {
			c.JSON(http.StatusMethodNotAllowed, gin.H{
				"error": "method not allowed",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
