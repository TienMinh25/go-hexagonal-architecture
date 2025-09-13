package domainauth

import (
	domainuser "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/user"
	"github.com/google/uuid"
)

// TokenPayload is an entity that represents the payload of the token
type TokenPayload struct {
	ID     uuid.UUID
	UserID uint64
	Role   domainuser.UserRole
}
