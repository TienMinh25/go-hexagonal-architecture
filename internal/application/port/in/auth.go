//go:generate mockgen -source=auth.go -destination=../../mock/auth.go -package=mock
package portin

import (
	"context"
)

// UserService is an interface for interacting with user authentication-related business logic
type AuthService interface {
	// Login authenticates a user by email and password and returns a token
	Login(ctx context.Context, email, password string) (string, error)
}
