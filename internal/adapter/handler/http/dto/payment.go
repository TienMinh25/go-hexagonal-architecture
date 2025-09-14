package dto

import domainpayment "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/payment"

// PaymentResponse represents a payment response body
type PaymentResponse struct {
	ID   uint64                    `json:"id" example:"1"`
	Name string                    `json:"name" example:"Tunai"`
	Type domainpayment.PaymentType `json:"type" example:"CASH"`
	Logo string                    `json:"logo" example:"https://example.com/cash.png"`
}

// CreatePaymentRequest represents a request body for creating a new payment
type CreatePaymentRequest struct {
	Name string                    `json:"name" binding:"required" example:"Tunai"`
	Type domainpayment.PaymentType `json:"type" binding:"required" example:"CASH"`
	Logo string                    `json:"logo" binding:"omitempty,required" example:"https://example.com/cash.png"`
}

// GetPaymentRequest represents a request body for retrieving a payment
type GetPaymentRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}

// ListPaymentsRequest represents a request body for listing payments
type ListPaymentsRequest struct {
	Skip  uint64 `form:"skip" binding:"required,min=0" example:"0"`
	Limit uint64 `form:"limit" binding:"required,min=5" example:"5"`
}

// UpdatePaymentRequest represents a request body for updating a payment
type UpdatePaymentRequest struct {
	Name string                    `json:"name" binding:"omitempty,required" example:"Gopay"`
	Type domainpayment.PaymentType `json:"type" binding:"omitempty,required,payment_type" example:"E-WALLET"`
	Logo string                    `json:"logo" binding:"omitempty,required" example:"https://example.com/gopay.png"`
}

// DeletePaymentRequest represents a request body for deleting a payment
type DeletePaymentRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}
