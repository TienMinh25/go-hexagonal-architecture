//go:generate mockgen -source=token.go -destination=../../mock/token.go -package=mock
package portout

import (
	domainauth "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/auth"
	domainuser "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/user"
)


// TokenService is an interface for interacting with token-related business logic
type TokenService interface {
	// CreateToken creates a new token for a given user
	CreateToken(user *domainuser.User) (string, error)
	// VerifyToken verifies the token and returns the payload
	VerifyToken(token string) (*domainauth.TokenPayload, error)
}
