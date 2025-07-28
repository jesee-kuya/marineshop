package main

import (
	"log"
	"net/http"

	"github.com/jesee-kuya/marineshop/pkg/db"
	"github.com/jesee-kuya/marineshop/pkg/handler"
	"github.com/jesee-kuya/marineshop/pkg/model"
	"github.com/jesee-kuya/marineshop/pkg/repository"
)

func main() {
	db, err := db.Init()
	if err != nil {
		log.Default().Println(err)
	}

	app := handler.App{
		Queries: repository.Query{
			Db: db,
		},
		User: &model.User{},
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: app.Routes(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Default().Println(err)
			return
		}
	}()

	log.Default().Printf("Server started on port %s", server.Addr)
	select {}
}
