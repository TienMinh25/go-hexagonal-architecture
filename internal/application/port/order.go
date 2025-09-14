package port

import (
	"context"

	domainorder "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/order"
)

// OrderRepository is an interface for interacting with order-related data
//
//go:generate mockgen -destination=../mock/order-repository.go -package=mock github.com/TienMinh25/go-hexagonal-architecture/internal/application/port OrderRepository
type OrderRepository interface {
	// CreateOrder inserts a new order into the database
	CreateOrder(ctx context.Context, order *domainorder.Order) (*domainorder.Order, error)
	// GetOrderByID selects an order by id
	GetOrderByID(ctx context.Context, id uint64) (*domainorder.Order, error)
	// ListOrders selects a list of orders with pagination
	ListOrders(ctx context.Context, skip, limit uint64) ([]domainorder.Order, error)
}

// OrderService is an interface for interacting with order-related business logic
//
//go:generate mockgen -destination=../mock/order-service.go -package=mock github.com/TienMinh25/go-hexagonal-architecture/internal/application/port OrderService
type OrderService interface {
	// CreateOrder creates a new order
	CreateOrder(ctx context.Context, order *domainorder.Order) (*domainorder.Order, error)
	// GetOrder returns an order by id
	GetOrder(ctx context.Context, id uint64) (*domainorder.Order, error)
	// ListOrders returns a list of orders with pagination
	ListOrders(ctx context.Context, skip, limit uint64) ([]domainorder.Order, error)
}
