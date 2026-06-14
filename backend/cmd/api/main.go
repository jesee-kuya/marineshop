package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jesee-kuya/marineshop/config"
	"github.com/jesee-kuya/marineshop/database"
	"github.com/jesee-kuya/marineshop/handler"
	"github.com/jesee-kuya/marineshop/middleware"
	"github.com/jesee-kuya/marineshop/repository"
	"github.com/jesee-kuya/marineshop/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, &cfg.JWT)
	middleware := middleware.NewMiddleware()

	shop := handler.Marineshop{
		AuthService: authService,
		Middleware:  middleware,
	}

	router := shop.SetupRoutes()

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	// Start server in goroutine for graceful shutdown
	go func() {
		if err := router.Run(addr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

}
