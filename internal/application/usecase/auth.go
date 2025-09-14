package usecase

import (
	"context"

	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/port"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/util"
)

/**
 * authUsecase implements port.AuthService interface
 * and provides an access to the user repository
 * and token service
 */
type authUsecase struct {
	repo port.UserRepository
	ts   port.TokenService
}

// NewAuthUsecase creates a new auth service instance
func NewAuthUsecase(repo port.UserRepository, ts port.TokenService) port.AuthService {
	return &authUsecase{
		repo: repo,
		ts:   ts,
	}
}

// Login gives a registered user an access token if the credentials are valid
func (as *authUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := as.repo.GetUserByEmail(ctx, email)

	if err != nil {
		if err == domain.ErrDataNotFound {
			return "", domain.ErrInvalidCredentials
		}
		return "", domain.ErrInternal
	}

	err = util.ComparePassword(password, user.Password)

	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	accessToken, err := as.ts.CreateToken(user)

	if err != nil {
		return "", domain.ErrTokenCreation
	}

	return accessToken, nil
}
