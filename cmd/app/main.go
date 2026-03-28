package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/Alexeyts0Y/test_task_em/internal/handlers"
	"github.com/Alexeyts0Y/test_task_em/internal/repository"
	"github.com/Alexeyts0Y/test_task_em/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	if err := repository.RunMigrations(dsn); err != nil {
		slog.Error("Migration error", "error", err)
		os.Exit(1)
	}

	pool, err := repository.InitDB(context.Background(), dsn)
	if err != nil {
		slog.Error("DB init error", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	repo := repository.NewSubscriptionRepo(pool)

	repository.RunMigrations(dsn)

	svc := service.NewSubscriptionService(repo)

	h := handlers.NewHandler(svc)

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/subscriptions", h.Create)
		api.GET("/subscriptions", h.List)
		api.GET("/subscriptions/:id", h.Get)
		api.PUT("/subscriptions/:id", h.Update)
		api.DELETE("/subscriptions/:id", h.Delete)
		api.GET("/subscriptions/cost", h.CalculateCost)
	}

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
