package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jesee-kuya/marineshop/pkg/model"
)

func (app *App) SignUp(w http.ResponseWriter, r *http.Request) {
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		app.JSONResponse(w, r, http.StatusBadRequest, err.Error(), Error)
		return
	}

	if err := user.ValidateUserDetails(); err != nil {
		app.JSONResponse(w, r, http.StatusBadRequest, err.Error(), Error)
		return
	}

	if err := app.Query.InsertData("users", []string{
		"email",
		"username",
		"password",
		"role",
		"status",
	}, []any{
		user.Email,
		user.Username,
		user.Password,
		user.Role,
		user.Status,
	}); err != nil {
		app.JSONResponse(w, r, http.StatusInternalServerError, err.Error(), Error)
		return
	}

	app.JSONResponse(w, r, http.StatusOK, "user created", Success)
}
