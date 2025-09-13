//go:generate mockgen -source=order.go -destination=../../mock/user-repository.go -package=mock
package portout

import (
	"context"

	domainuser "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/user"
)

// UserRepository is an interface for interacting with user-related data
type UserRepository interface {
	// CreateUser inserts a new user into the database
	CreateUser(ctx context.Context, user *domainuser.User) (*domainuser.User, error)
	// GetUserByID selects a user by id
	GetUserByID(ctx context.Context, id uint64) (*domainuser.User, error)
	// GetUserByEmail selects a user by email
	GetUserByEmail(ctx context.Context, email string) (*domainuser.User, error)
	// ListUsers selects a list of users with pagination
	ListUsers(ctx context.Context, skip, limit uint64) ([]domainuser.User, error)
	// UpdateUser updates a user
	UpdateUser(ctx context.Context, user *domainuser.User) (*domainuser.User, error)
	// DeleteUser deletes a user
	DeleteUser(ctx context.Context, id uint64) error
}
