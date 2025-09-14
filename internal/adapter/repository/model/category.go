package model

import (
	"time"

	domaincategory "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/category"
)

type Category struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (c Category) ToDomain() *domaincategory.Category {
	return &domaincategory.Category{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
