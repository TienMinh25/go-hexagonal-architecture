package modelv1

import (
	"time"

	domainuser "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/user"
)

// UserResponse represents a user response body
type UserResponse struct {
	ID        uint64    `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"test@example.com"`
	CreatedAt time.Time `json:"created_at" example:"1970-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"1970-01-01T00:00:00Z"`
}

// RegisterRequest represents the request body for creating a user
type RegisterRequest struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"test@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"12345678"`
}

// ListUsersRequest represents the request body for listing users
type ListUsersRequest struct {
	Skip  uint64 `form:"skip" binding:"required,min=0" example:"0"`
	Limit uint64 `form:"limit" binding:"required,min=5" example:"5"`
}

// GetUserRequest represents the request body for getting a user
type GetUserRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	Name     string              `json:"name" binding:"omitempty,required" example:"John Doe"`
	Email    string              `json:"email" binding:"omitempty,required,email" example:"test@example.com"`
	Password string              `json:"password" binding:"omitempty,required,min=8" example:"12345678"`
	Role     domainuser.UserRole `json:"role" binding:"omitempty,required,user_role" example:"admin"`
}

// DeleteUserRequest represents the request body for deleting a user
type DeleteUserRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}
