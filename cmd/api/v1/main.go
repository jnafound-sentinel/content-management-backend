package main

import (
	"fmt"
	"log"

	"jna-manager/internal/config"
	"jna-manager/internal/handler"
	"jna-manager/internal/repository/database"
	"jna-manager/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repositories
	userRepo := database.NewUserRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)

	emailService, err := service.NewEmailService(cfg)
	if err != nil {
		log.Fatal("Could not load Email Service configurations and Templates.")
	}

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService, cfg)
	authHandler := handler.NewAuthHandler(userService, emailService, cfg)

	// Setup router
	r := gin.Default()

	api := r.Group("/api")
	{
		userHandler.RegisterRoutes(api)
		authHandler.RegisterRoutes(api)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Port))
}
