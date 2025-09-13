//go:generate mockgen -source=category.go -destination=../../mock/category-repository.go -package=mock
package portout

import (
	"context"

	domaincategory "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/category"
)

// CategoryRepository is an interface for interacting with category-related data
type CategoryRepository interface {
	// CreateCategory inserts a new category into the database
	CreateCategory(ctx context.Context, category *domaincategory.Category) (*domaincategory.Category, error)
	// GetCategoryByID selects a category by id
	GetCategoryByID(ctx context.Context, id uint64) (*domaincategory.Category, error)
	// ListCategories selects a list of categories with pagination
	ListCategories(ctx context.Context, skip, limit uint64) ([]domaincategory.Category, error)
	// UpdateCategory updates a category
	UpdateCategory(ctx context.Context, category *domaincategory.Category) (*domaincategory.Category, error)
	// DeleteCategory deletes a category
	DeleteCategory(ctx context.Context, id uint64) error
}
