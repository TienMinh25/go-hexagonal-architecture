//go:generate mockgen -source=product.go -destination=../../mock/product-repository.go -package=mock
package portout

import (
	"context"

	domainproduct "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/product"
)

// ProductRepository is an interface for interacting with product-related data
type ProductRepository interface {
	// CreateProduct inserts a new product into the database
	CreateProduct(ctx context.Context, product *domainproduct.Product) (*domainproduct.Product, error)
	// GetProductByID selects a product by id
	GetProductByID(ctx context.Context, id uint64) (*domainproduct.Product, error)
	// ListProducts selects a list of products with pagination
	ListProducts(ctx context.Context, search string, categoryId, skip, limit uint64) ([]domainproduct.Product, error)
	// UpdateProduct updates a product
	UpdateProduct(ctx context.Context, product *domainproduct.Product) (*domainproduct.Product, error)
	// DeleteProduct deletes a product
	DeleteProduct(ctx context.Context, id uint64) error
}
