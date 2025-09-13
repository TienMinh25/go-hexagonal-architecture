//go:generate mockgen -source=order.go -destination=../../mock/order-service.go -package=mock
package portin

import (
	"context"

	domainorder "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/order"
)

// OrderService is an interface for interacting with order-related business logic
type OrderService interface {
	// CreateOrder creates a new order
	CreateOrder(ctx context.Context, order *domainorder.Order) (*domainorder.Order, error)
	// GetOrder returns an order by id
	GetOrder(ctx context.Context, id uint64) (*domainorder.Order, error)
	// ListOrders returns a list of orders with pagination
	ListOrders(ctx context.Context, skip, limit uint64) ([]domainorder.Order, error)
}
