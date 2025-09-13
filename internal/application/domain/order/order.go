package domainorder

import (
	"time"

	domainpayment "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/payment"
	domainuser "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/user"
	"github.com/google/uuid"
)

// Order is an entity that represents an order
type Order struct {
	ID           uint64
	UserID       uint64
	PaymentID    uint64
	CustomerName string
	TotalPrice   float64
	TotalPaid    float64
	TotalReturn  float64
	ReceiptCode  uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         *domainuser.User
	Payment      *domainpayment.Payment
	Products     []OrderProduct
}
