package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/TienMinh25/go-hexagonal-architecture/docs"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/adapter/auth/paseto"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/adapter/config"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/adapter/handler/http"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/adapter/logger"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/adapter/storage/postgres"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/adapter/storage/postgres/repository"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/adapter/storage/redis"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/usecase"
)

//	@title						Go Sale API demo for hexagonal architecture
//	@version					1.0
//	@description				This is a simple RESTful Point of Sale (POS) Service API written in Go using Gin web framework, PostgreSQL database, and Redis cache.
//
//	@contact.name				Tien Minh
//	@contact.url				https://github.com/TienMinh25/go-hexagonal-architecture
//	@contact.email				letienminh2512@gmail.com
//
//	@license.name				APACHE
//	@license.url				https://github.com/TienMinh25/go-hexagonal-architecture/blob/main/LICENSE
//
//	@host						localhost
//	@BasePath					/v1
//	@schemes					http https
//
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and the access token.
func main() {
	// Load environment variables
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	// Set logger
	logger.Set(config.App)

	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Init database
	ctx := context.Background()
	db, err := postgres.New(ctx, config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Successfully connected to the database", "db", config.DB.Connection)

	// Migrate database
	err = db.Migrate()
	if err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully migrated the database")

	// Init cache service
	cache, err := redis.New(ctx, config.Redis)
	if err != nil {
		slog.Error("Error initializing cache connection", "error", err)
		os.Exit(1)
	}
	defer cache.Close()

	slog.Info("Successfully connected to the cache server")

	// Init token service
	token, err := paseto.New(config.Token)
	if err != nil {
		slog.Error("Error initializing token service", "error", err)
		os.Exit(1)
	}

	// Dependency injection
	// User
	userRepo := repository.NewUserRepository(db)
	userService := usecase.NewUserUsecase(userRepo, cache)
	userHandler := http.NewUserHandler(userService)

	// Auth
	authService := usecase.NewAuthUsecase(userRepo, token)
	authHandler := http.NewAuthHandler(authService)

	// Payment
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := usecase.NewPaymentUsecase(paymentRepo, cache)
	paymentHandler := http.NewPaymentHandler(paymentService)

	// Category
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := usecase.NewCategoryUsecase(categoryRepo, cache)
	categoryHandler := http.NewCategoryHandler(categoryService)

	// Product
	productRepo := repository.NewProductRepository(db)
	productService := usecase.NewProductUsecase(productRepo, categoryRepo, cache)
	productHandler := http.NewProductHandler(productService)

	// Order
	orderRepo := repository.NewOrderRepository(db)
	orderService := usecase.NewOrderUsecase(orderRepo, productRepo, categoryRepo, userRepo, paymentRepo, cache)
	orderHandler := http.NewOrderHandler(orderService)

	// Init router
	router, err := http.NewRouter(
		config.HTTP,
		token,
		*userHandler,
		*authHandler,
		*paymentHandler,
		*categoryHandler,
		*productHandler,
		*orderHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}