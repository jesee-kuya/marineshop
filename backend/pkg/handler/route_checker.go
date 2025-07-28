package handler

import (
	"net/http"
	"strings"
)

func (app *App) RouteChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			allowedURL, ok := AllowedRoutes[r.URL.Path]
			if !ok && !strings.HasPrefix(r.URL.Path, "/pkg/db/media/") {
				app.JSONResponse(w, r, http.StatusNotFound, "route not found", Error)
				return
			}

			method_found := false

			for _, method := range allowedURL {
				if r.Method == method {
					method_found = true
				}
			}

			if strings.HasPrefix(r.URL.Path, "/pkg/db/media/") && r.Method == http.MethodGet {
				method_found = true
			}

			if !method_found {
				app.JSONResponse(w, r, http.StatusMethodNotAllowed, "method not allowed", Error)
				return
			}
			next.ServeHTTP(w, r)
		},
	)
}
