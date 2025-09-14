package modelv1

import "time"

// ProductResponse represents a product response body
type ProductResponse struct {
	ID        uint64           `json:"id" example:"1"`
	SKU       string           `json:"sku" example:"9a4c25d3-9786-492c-b084-85cb75c1ee3e"`
	Name      string           `json:"name" example:"Chiki Ball"`
	Stock     int64            `json:"stock" example:"100"`
	Price     float64          `json:"price" example:"5000"`
	Image     string           `json:"image" example:"https://example.com/chiki-ball.png"`
	Category  CategoryResponse `json:"category"`
	CreatedAt time.Time        `json:"created_at" example:"1970-01-01T00:00:00Z"`
	UpdatedAt time.Time        `json:"updated_at" example:"1970-01-01T00:00:00Z"`
}

// CreateProductRequest represents a request body for creating a new product
type CreateProductRequest struct {
	CategoryID uint64  `json:"category_id" binding:"required,min=1" example:"1"`
	Name       string  `json:"name" binding:"required" example:"Chiki Ball"`
	Image      string  `json:"image" binding:"required" example:"https://example.com/chiki-ball.png"`
	Price      float64 `json:"price" binding:"required,min=0" example:"5000"`
	Stock      int64   `json:"stock" binding:"required,min=0" example:"100"`
}

// GetProductRequest represents a request body for retrieving a product
type GetProductRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}

// ListProductsRequest represents a request body for listing products
type ListProductsRequest struct {
	CategoryID uint64 `form:"category_id" binding:"omitempty,min=1" example:"1"`
	Query      string `form:"q" binding:"omitempty" example:"Chiki"`
	Skip       uint64 `form:"skip" binding:"required,min=0" example:"0"`
	Limit      uint64 `form:"limit" binding:"required,min=5" example:"5"`
}

// UpdateProductRequest represents a request body for updating a product
type UpdateProductRequest struct {
	CategoryID uint64  `json:"category_id" binding:"omitempty,required,min=1" example:"1"`
	Name       string  `json:"name" binding:"omitempty,required" example:"Nutrisari Jeruk"`
	Image      string  `json:"image" binding:"omitempty,required" example:"https://example.com/nutrisari-jeruk.png"`
	Price      float64 `json:"price" binding:"omitempty,required,min=0" example:"2000"`
	Stock      int64   `json:"stock" binding:"omitempty,required,min=0" example:"200"`
}

// DeleteProductRequest represents a request body for deleting a product
type DeleteProductRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}
