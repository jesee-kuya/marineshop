package handler

import (
	"net/http"

	"github.com/jesee-kuya/marineshop/pkg/model"
	"github.com/jesee-kuya/marineshop/pkg/repository"
)

var AllowedRoutes = map[string][]string{}

type App struct {
	Query *repository.Query
	User  *model.User
}

func (app *App) Routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("pkg/db/media"))
	mux.Handle("/pkg/db/media/", http.StripPrefix("/pkg/db/media/", fs))

	return mux
}
