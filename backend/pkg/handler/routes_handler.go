package handler

import (
	"net/http"

	"github.com/jesee-kuya/marineshop/pkg/model"
	"github.com/jesee-kuya/marineshop/pkg/repository"
)

var AllowedRoutes = map[string][]string{
	"/api/signup": {"POST", "OPTIONS"},
}

type App struct {
	Queries repository.Query
	User    *model.User
}

func (app *App) Routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("pkg/db/media"))
	mux.Handle("/pkg/db/media/", http.StripPrefix("/pkg/db/media/", fs))

	mux.Handle("/api/signup", app.RouteChecker(http.HandlerFunc(app.SignUp)))

	return mux
}
