package dto

import "time"

// OrderResponse represents an order response body
type OrderResponse struct {
	ID           uint64                 `json:"id" example:"1"`
	UserID       uint64                 `json:"user_id" example:"1"`
	PaymentID    uint64                 `json:"payment_type_id" example:"1"`
	CustomerName string                 `json:"customer_name" example:"John Doe"`
	TotalPrice   float64                `json:"total_price" example:"100000"`
	TotalPaid    float64                `json:"total_paid" example:"100000"`
	TotalReturn  float64                `json:"total_return" example:"0"`
	ReceiptCode  string                 `json:"receipt_id" example:"4979cf6e-d215-4ff8-9d0d-b3e99bcc7750"`
	Products     []OrderProductResponse `json:"products"`
	PaymentType  PaymentResponse        `json:"payment_type"`
	CreatedAt    time.Time              `json:"created_at" example:"1970-01-01T00:00:00Z"`
	UpdatedAt    time.Time              `json:"updated_at" example:"1970-01-01T00:00:00Z"`
}

// orderProductResponse represents an order product response body
type OrderProductResponse struct {
	ID               uint64          `json:"id" example:"1"`
	OrderID          uint64          `json:"order_id" example:"1"`
	ProductID        uint64          `json:"product_id" example:"1"`
	Quantity         int64           `json:"qty" example:"1"`
	Price            float64         `json:"price" example:"100000"`
	TotalNormalPrice float64         `json:"total_normal_price" example:"100000"`
	TotalFinalPrice  float64         `json:"total_final_price" example:"100000"`
	Product          ProductResponse `json:"product"`
	CreatedAt        time.Time       `json:"created_at" example:"1970-01-01T00:00:00Z"`
	UpdatedAt        time.Time       `json:"updated_at" example:"1970-01-01T00:00:00Z"`
}

// OrderProductRequest represents an order product request body
type OrderProductRequest struct {
	ProductID uint64 `json:"product_id" binding:"required,min=1" example:"1"`
	Quantity  int64  `json:"qty" binding:"required,number" example:"1"`
}

// CreateOrderRequest represents a request body for creating a new order
type CreateOrderRequest struct {
	PaymentID    uint64                `json:"payment_id" binding:"required" example:"1"`
	CustomerName string                `json:"customer_name" binding:"required" example:"John Doe"`
	TotalPaid    int64                 `json:"total_paid" binding:"required" example:"100000"`
	Products     []OrderProductRequest `json:"products" binding:"required"`
}

// GetOrderRequest represents a request body for retrieving an order
type GetOrderRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}

// ListOrdersRequest represents a request body for listing orders
type ListOrdersRequest struct {
	Skip  uint64 `form:"skip" binding:"required,min=0" example:"0"`
	Limit uint64 `form:"limit" binding:"required,min=5" example:"5"`
}
