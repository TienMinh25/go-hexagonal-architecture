//go:generate mockgen -source=payment.go -destination=../../mock/payment-service.go -package=mock
package portin

import (
	"context"

	domainpayment "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/payment"
)

// PaymentService is an interface for interacting with payment-related business logic
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
