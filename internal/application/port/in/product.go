//go:generate mockgen -source=product.go -destination=../../mock/product-service.go -package=mock
package portin

import (
	"context"

	domainproduct "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/product"
)

// ProductService is an interface for interacting with product-related business logic
type ProductService interface {
	// CreateProduct creates a new product
	CreateProduct(ctx context.Context, product *domainproduct.Product) (*domainproduct.Product, error)
	// GetProduct returns a product by id
	GetProduct(ctx context.Context, id uint64) (*domainproduct.Product, error)
	// ListProducts returns a list of products with pagination
	ListProducts(ctx context.Context, search string, categoryId, skip, limit uint64) ([]domainproduct.Product, error)
	// UpdateProduct updates a product
	UpdateProduct(ctx context.Context, product *domainproduct.Product) (*domainproduct.Product, error)
	// DeleteProduct deletes a product
	DeleteProduct(ctx context.Context, id uint64) error
}
