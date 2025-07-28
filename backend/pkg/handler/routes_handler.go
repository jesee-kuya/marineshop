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

	return mux
}
