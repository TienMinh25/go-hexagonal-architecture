//go:generate mockgen -source=payment.go -destination=../../mock/payment-repository.go -package=mock
package portout

import (
	"context"

	domainpayment "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/payment"
)

// PaymentRepository is an interface for interacting with payment-related data
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
