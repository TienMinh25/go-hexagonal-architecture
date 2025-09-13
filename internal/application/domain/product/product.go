package domainproduct

import (
	"time"

	domaincategory "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/category"
	"github.com/google/uuid"
)

// Product is an entity that represents a product
type Product struct {
	ID         uint64
	CategoryID uint64
	SKU        uuid.UUID
	Name       string
	Stock      int64
	Price      float64
	Image      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Category   *domaincategory.Category
}
