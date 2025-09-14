package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         uint64    `db:"id"`
	CategoryID uint64    `db:"category_id"`
	SKU        uuid.UUID `db:"sku"`
	Name       string    `db:"name"`
	Stock      int64     `db:"stock"`
	Price      float64   `db:"price"`
	Image      string    `db:"image"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Category   *Category `db:"category"`
}
