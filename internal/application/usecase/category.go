package usecase

import (
	"context"

	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain"
	domaincategory "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/category"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/port"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/util"
)

/**
 * categoryUsecase implements port.CategoryService interface
 * and provides an access to the category repository
 * and cache service
 */
type categoryUsecase struct {
	repo  port.CategoryRepository
	cache port.CacheRepository
}

// NewCategoryUsecase creates a new category service instance
func NewCategoryUsecase(repo port.CategoryRepository, cache port.CacheRepository) *categoryUsecase {
	return &categoryUsecase{
		repo,
		cache,
	}
}

// CreateCategory creates a new category
func (cs *categoryUsecase) CreateCategory(ctx context.Context, category *domaincategory.Category) (*domaincategory.Category, error) {
	category, err := cs.repo.CreateCategory(ctx, category)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	cacheKey := util.GenerateCacheKey("category", category.ID)
	categorySerialized, err := util.Serialize(category)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = cs.cache.Set(ctx, cacheKey, categorySerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = cs.cache.DeleteByPrefix(ctx, "categories:*")
	if err != nil {
		return nil, domain.ErrInternal
	}

	return category, nil
}

// GetCategory retrieves a category by id
func (cs *categoryUsecase) GetCategory(ctx context.Context, id uint64) (*domaincategory.Category, error) {
	var category *domaincategory.Category

	cacheKey := util.GenerateCacheKey("category", id)
	cachedCategory, err := cs.cache.Get(ctx, cacheKey)
	if err == nil {
		err := util.Deserialize(cachedCategory, &category)
		if err != nil {
			return nil, domain.ErrInternal
		}
		return category, nil
	}

	category, err = cs.repo.GetCategoryByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	categorySerialized, err := util.Serialize(category)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = cs.cache.Set(ctx, cacheKey, categorySerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return category, nil
}

// ListCategories retrieves a list of categories
func (cs *categoryUsecase) ListCategories(ctx context.Context, skip, limit uint64) ([]domaincategory.Category, error) {
	var categories []domaincategory.Category

	params := util.GenerateCacheKeyParams(skip, limit)
	cacheKey := util.GenerateCacheKey("categories", params)

	cachedCategories, err := cs.cache.Get(ctx, cacheKey)
	if err == nil {
		err := util.Deserialize(cachedCategories, &categories)
		if err != nil {
			return nil, domain.ErrInternal
		}

		return categories, nil
	}

	categories, err = cs.repo.ListCategories(ctx, skip, limit)
	if err != nil {
		return nil, domain.ErrInternal
	}

	categoriesSerialized, err := util.Serialize(categories)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = cs.cache.Set(ctx, cacheKey, categoriesSerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return categories, nil
}

// UpdateCategory updates a category
func (cs *categoryUsecase) UpdateCategory(ctx context.Context, category *domaincategory.Category) (*domaincategory.Category, error) {
	existingCategory, err := cs.repo.GetCategoryByID(ctx, category.ID)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	emptyData := category.Name == ""
	sameData := existingCategory.Name == category.Name
	if emptyData || sameData {
		return nil, domain.ErrNoUpdatedData
	}

	_, err = cs.repo.UpdateCategory(ctx, category)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	cacheKey := util.GenerateCacheKey("category", category.ID)

	err = cs.cache.Delete(ctx, cacheKey)
	if err != nil {
		return nil, domain.ErrInternal
	}

	categorySerialized, err := util.Serialize(category)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = cs.cache.Set(ctx, cacheKey, categorySerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = cs.cache.DeleteByPrefix(ctx, "categories:*")
	if err != nil {
		return nil, domain.ErrInternal
	}

	return category, nil
}

// DeleteCategory deletes a category
func (cs *categoryUsecase) DeleteCategory(ctx context.Context, id uint64) error {
	_, err := cs.repo.GetCategoryByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return err
		}
		return domain.ErrInternal
	}

	cacheKey := util.GenerateCacheKey("category", id)

	err = cs.cache.Delete(ctx, cacheKey)
	if err != nil {
		return domain.ErrInternal
	}

	err = cs.cache.DeleteByPrefix(ctx, "categories:*")
	if err != nil {
		return domain.ErrInternal
	}

	return cs.repo.DeleteCategory(ctx, id)
}
