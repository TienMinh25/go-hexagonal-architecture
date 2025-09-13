package usecase

import (
	"context"

	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain"
	domainuser "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/user"
	portin "github.com/TienMinh25/go-hexagonal-architecture/internal/application/port/in"
	portout "github.com/TienMinh25/go-hexagonal-architecture/internal/application/port/out"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/util"
)

/**
 * userUsecase implements portin.UserService interface
 * and provides an access to the user repository
 * and cache service
 */
type userUsecase struct {
	repo  portout.UserRepository
	cache portout.CacheRepository
}

// NewUserUsecase creates a new user service instance
func NewUserUsecase(repo portout.UserRepository, cache portout.CacheRepository) portin.UserService {
	return &userUsecase{
		repo,
		cache,
	}
}

// Register creates a new user
func (us *userUsecase) Register(ctx context.Context, user *domainuser.User) (*domainuser.User, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, domain.ErrInternal
	}

	user.Password = hashedPassword

	user, err = us.repo.CreateUser(ctx, user)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	cacheKey := util.GenerateCacheKey("user", user.ID)
	userSerialized, err := util.Serialize(user)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = us.cache.Set(ctx, cacheKey, userSerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = us.cache.DeleteByPrefix(ctx, "users:*")
	if err != nil {
		return nil, domain.ErrInternal
	}

	return user, nil
}

// GetUser gets a user by ID
func (us *userUsecase) GetUser(ctx context.Context, id uint64) (*domainuser.User, error) {
	var user *domainuser.User

	cacheKey := util.GenerateCacheKey("user", id)
	cachedUser, err := us.cache.Get(ctx, cacheKey)
	if err == nil {
		err := util.Deserialize(cachedUser, &user)
		if err != nil {
			return nil, domain.ErrInternal
		}
		return user, nil
	}

	user, err = us.repo.GetUserByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	userSerialized, err := util.Serialize(user)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = us.cache.Set(ctx, cacheKey, userSerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return user, nil
}

// ListUsers lists all users
func (us *userUsecase) ListUsers(ctx context.Context, skip, limit uint64) ([]domainuser.User, error) {
	var users []domainuser.User

	params := util.GenerateCacheKeyParams(skip, limit)
	cacheKey := util.GenerateCacheKey("users", params)

	cachedUsers, err := us.cache.Get(ctx, cacheKey)
	if err == nil {
		err := util.Deserialize(cachedUsers, &users)
		if err != nil {
			return nil, domain.ErrInternal
		}
		return users, nil
	}

	users, err = us.repo.ListUsers(ctx, skip, limit)
	if err != nil {
		return nil, domain.ErrInternal
	}

	usersSerialized, err := util.Serialize(users)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = us.cache.Set(ctx, cacheKey, usersSerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return users, nil
}

// UpdateUser updates a user's name, email, and password
func (us *userUsecase) UpdateUser(ctx context.Context, user *domainuser.User) (*domainuser.User, error) {
	existingUser, err := us.repo.GetUserByID(ctx, user.ID)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	emptyData := user.Name == "" &&
		user.Email == "" &&
		user.Password == "" &&
		user.Role == ""
	sameData := existingUser.Name == user.Name &&
		existingUser.Email == user.Email &&
		existingUser.Role == user.Role
	if emptyData || sameData {
		return nil, domain.ErrNoUpdatedData
	}

	var hashedPassword string

	if user.Password != "" {
		hashedPassword, err = util.HashPassword(user.Password)
		if err != nil {
			return nil, domain.ErrInternal
		}
	}

	user.Password = hashedPassword

	_, err = us.repo.UpdateUser(ctx, user)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	cacheKey := util.GenerateCacheKey("user", user.ID)

	err = us.cache.Delete(ctx, cacheKey)
	if err != nil {
		return nil, domain.ErrInternal
	}

	userSerialized, err := util.Serialize(user)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = us.cache.Set(ctx, cacheKey, userSerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = us.cache.DeleteByPrefix(ctx, "users:*")
	if err != nil {
		return nil, domain.ErrInternal
	}

	return user, nil
}

// DeleteUser deletes a user by ID
func (us *userUsecase) DeleteUser(ctx context.Context, id uint64) error {
	_, err := us.repo.GetUserByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return err
		}
		return domain.ErrInternal
	}

	cacheKey := util.GenerateCacheKey("user", id)

	err = us.cache.Delete(ctx, cacheKey)
	if err != nil {
		return domain.ErrInternal
	}

	err = us.cache.DeleteByPrefix(ctx, "users:*")
	if err != nil {
		return domain.ErrInternal
	}

	return us.repo.DeleteUser(ctx, id)
}
