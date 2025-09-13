package http

import (
	domainpayment "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/payment"
	domainuser "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/user"
	"github.com/go-playground/validator/v10"
)

// userRoleValidator is a custom validator for validating user roles
var userRoleValidator validator.Func = func(fl validator.FieldLevel) bool {
	userRole := fl.Field().Interface().(domainuser.UserRole)

	switch userRole {
	case "admin", "cashier":
		return true
	default:
		return false
	}
}

// paymentTypeValidator is a custom validator for validating payment types
var paymentTypeValidator validator.Func = func(fl validator.FieldLevel) bool {
	paymentType := fl.Field().Interface().(domainpayment.PaymentType)

	switch paymentType {
	case "CASH", "E-WALLET", "EDC":
		return true
	default:
		return false
	}
}