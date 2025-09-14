package port

import (
	"context"

	domainpayment "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/payment"
)

// PaymentRepository is an interface for interacting with payment-related data
//
//go:generate mockgen -destination=../mock/payment-repository.go -package=mock github.com/TienMinh25/go-hexagonal-architecture/internal/application/port PaymentRepository
type PaymentRepository interface {
	// CreatePayment inserts a new payment into the database
	CreatePayment(ctx context.Context, payment *domainpayment.Payment) (*domainpayment.Payment, error)
	// GetPaymentByID selects a payment by id
	GetPaymentByID(ctx context.Context, id uint64) (*domainpayment.Payment, error)
	// ListPayments selects a list of payments with pagination
	ListPayments(ctx context.Context, skip, limit uint64) ([]domainpayment.Payment, error)
	// UpdatePayment updates a payment
	UpdatePayment(ctx context.Context, payment *domainpayment.Payment) (*domainpayment.Payment, error)
	// DeletePayment deletes a payment
	DeletePayment(ctx context.Context, id uint64) error
}

// PaymentService is an interface for interacting with payment-related business logic
//
//go:generate mockgen -destination=../mock/payment-service.go -package=mock github.com/TienMinh25/go-hexagonal-architecture/internal/application/port PaymentService
type PaymentService interface {
	// CreatePayment creates a new payment
	CreatePayment(ctx context.Context, payment *domainpayment.Payment) (*domainpayment.Payment, error)
	// GetPayment returns a payment by id
	GetPayment(ctx context.Context, id uint64) (*domainpayment.Payment, error)
	// ListPayments returns a list of payments with pagination
	ListPayments(ctx context.Context, skip, limit uint64) ([]domainpayment.Payment, error)
	// UpdatePayment updates a payment
	UpdatePayment(ctx context.Context, payment *domainpayment.Payment) (*domainpayment.Payment, error)
	// DeletePayment deletes a payment
	DeletePayment(ctx context.Context, id uint64) error
}
