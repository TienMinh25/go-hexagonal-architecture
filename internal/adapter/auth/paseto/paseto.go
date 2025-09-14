package paseto

import (
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/TienMinh25/go-hexagonal-architecture/config"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain"
	domainauth "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/auth"
	domainuser "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/user"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/port"
	"github.com/google/uuid"
)

/**
 * pasetoToken implements portout.TokenService interface
 * and provides an access to the paseto library
 */
type pasetoToken struct {
	token    *paseto.Token
	key      *paseto.V4SymmetricKey
	parser   *paseto.Parser
	duration time.Duration
}

// New creates a new paseto instance
func New(config *config.Token) (port.TokenService, error) {
	durationStr := config.Duration
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return nil, domain.ErrTokenDuration
	}

	token := paseto.NewToken()
	key := paseto.NewV4SymmetricKey()
	parser := paseto.NewParser()

	return &pasetoToken{
		&token,
		&key,
		&parser,
		duration,
	}, nil
}

// CreateToken creates a new paseto token
func (pt *pasetoToken) CreateToken(user *domainuser.User) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", domain.ErrTokenCreation
	}

	payload := &domainauth.TokenPayload{
		ID:     id,
		UserID: user.ID,
		Role:   user.Role,
	}

	err = pt.token.Set("payload", payload)
	if err != nil {
		return "", domain.ErrTokenCreation
	}

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(pt.duration)

	pt.token.SetIssuedAt(issuedAt)
	pt.token.SetNotBefore(issuedAt)
	pt.token.SetExpiration(expiredAt)

	token := pt.token.V4Encrypt(*pt.key, nil)

	return token, nil
}

// VerifyToken verifies the paseto token
func (pt *pasetoToken) VerifyToken(token string) (*domainauth.TokenPayload, error) {
	var payload *domainauth.TokenPayload

	parsedToken, err := pt.parser.ParseV4Local(*pt.key, token, nil)
	if err != nil {
		if err.Error() == "this token has expired" {
			return nil, domain.ErrExpiredToken
		}
		return nil, domain.ErrInvalidToken
	}

	err = parsedToken.Get("payload", &payload)
	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	return payload, nil
}
