package main

import (
	"fmt"
	"log"

	"jna-manager/internal/config"
	"jna-manager/internal/handler"
	"jna-manager/internal/migrations"
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

	result := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	if result.Error != nil {
		log.Fatal("failed to enable uuid-ossp extension: %w", result.Error)
	}

	err = db.AutoMigrate(migrations.GetMigrationModels()...)
	if err != nil {
		log.Fatal("Failed to perform Database Migrations")
	}

	userRepo := database.NewUserRepository(db)
	blogRepo := database.NewBlogRepository(db)

	userService := service.NewUserService(userRepo)
	blogService := service.NewBlogService(blogRepo)

	emailService, err := service.NewEmailService(cfg)
	if err != nil {
		log.Fatal("Could not load Email Service configurations and Templates.")
	}

	userHandler := handler.NewUserHandler(userService, cfg)
	authHandler := handler.NewAuthHandler(userService, emailService, cfg)
	blogHandler := handler.NewBlogHandler(blogService, cfg)

	r := gin.Default()

	api := r.Group("/api")
	{
		userHandler.RegisterRoutes(api)
		authHandler.RegisterRoutes(api)
		blogHandler.RegisterRoutes(api)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Port))
}
