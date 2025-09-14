package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID           uint64         `db:"id"`
	UserID       uint64         `db:"user_id"`
	PaymentID    uint64         `db:"payment_id"`
	CustomerName string         `db:"customer_name"`
	TotalPrice   float64        `db:"total_price"`
	TotalPaid    float64        `db:"total_paid"`
	TotalReturn  float64        `db:"total_return"`
	ReceiptCode  uuid.UUID      `db:"receipt_code"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
	User         *User          `db:"user"`
	Payment      *Payment       `db:"payment"`
	Products     []OrderProduct `db:"products"`
}

type OrderProduct struct {
	ID         uint64
	OrderID    uint64
	ProductID  uint64
	Quantity   int64
	TotalPrice float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Order      *Order
	Product    *Product
}
