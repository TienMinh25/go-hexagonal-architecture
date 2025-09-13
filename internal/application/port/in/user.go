//go:generate mockgen -source=user.go -destination=../../mock/user-service.go -package=mock
package portin

import (
	"context"

	domainuser "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/user"
)

// UserService is an interface for interacting with user-related business logic
type UserService interface {
	// Register registers a new user
	Register(ctx context.Context, user *domainuser.User) (*domainuser.User, error)
	// GetUser returns a user by id
	GetUser(ctx context.Context, id uint64) (*domainuser.User, error)
	// ListUsers returns a list of users with pagination
	ListUsers(ctx context.Context, skip, limit uint64) ([]domainuser.User, error)
	// UpdateUser updates a user
	UpdateUser(ctx context.Context, user *domainuser.User) (*domainuser.User, error)
	// DeleteUser deletes a user
	DeleteUser(ctx context.Context, id uint64) error
}
