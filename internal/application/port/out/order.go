//go:generate mockgen -source=order.go -destination=../../mock/order-repository.go -package=mock
package portout

import (
	"context"

	domainorder "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/order"
)

// OrderRepository is an interface for interacting with order-related data
type OrderRepository interface {
	// CreateOrder inserts a new order into the database
	CreateOrder(ctx context.Context, order *domainorder.Order) (*domainorder.Order, error)
	// GetOrderByID selects an order by id
	GetOrderByID(ctx context.Context, id uint64) (*domainorder.Order, error)
	// ListOrders selects a list of orders with pagination
	ListOrders(ctx context.Context, skip, limit uint64) ([]domainorder.Order, error)
}
