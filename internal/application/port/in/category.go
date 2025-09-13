//go:generate mockgen -source=category.go -destination=../../mock/category-service.go -package=mock
package portin

import (
	"context"

	domaincategory "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/category"
)

// CategoryService is an interface for interacting with category-related business logic
type CategoryService interface {
	// CreateCategory creates a new category
	CreateCategory(ctx context.Context, category *domaincategory.Category) (*domaincategory.Category, error)
	// GetCategory returns a category by id
	GetCategory(ctx context.Context, id uint64) (*domaincategory.Category, error)
	// ListCategories returns a list of categories with pagination
	ListCategories(ctx context.Context, skip, limit uint64) ([]domaincategory.Category, error)
	// UpdateCategory updates a category
	UpdateCategory(ctx context.Context, category *domaincategory.Category) (*domaincategory.Category, error)
	// DeleteCategory deletes a category
	DeleteCategory(ctx context.Context, id uint64) error
}