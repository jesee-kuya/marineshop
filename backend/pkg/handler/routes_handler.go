package handler

import (
	"net/http"

	"github.com/jesee-kuya/marineshop/pkg/model"
)

var AllowedRoutes = map[string][]string{}

type App struct {
	User *model.User
}

func (app *App) Routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("pkg/db/media"))
	mux.Handle("/pkg/db/media/", http.StripPrefix("/pkg/db/media/", fs))

	return mux
}
