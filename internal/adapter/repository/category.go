package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	storagepostgres "github.com/TienMinh25/go-hexagonal-architecture/infrastructure/storage/postgres"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/adapter/repository/model"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain"
	domaincategory "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/category"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/port"
	"github.com/jackc/pgx/v5"
)

/**
 * categoryRepository implements port.CategoryRepository interface
 * and provides an access to the postgres database
 */
type categoryRepository struct {
	db *storagepostgres.DB
}

// NewCategoryRepository creates a new category repository instance
func NewCategoryRepository(db *storagepostgres.DB) port.CategoryRepository {
	return &categoryRepository{
		db,
	}
}

// CreateCategory creates a new category record in the database
func (cr *categoryRepository) CreateCategory(ctx context.Context, category *domaincategory.Category) (*domaincategory.Category, error) {
	query := cr.db.QueryBuilder.Insert("categories").
		Columns("name").
		Values(category.Name).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var returnCategory model.Category

	err = cr.db.QueryRow(ctx, sql, args...).Scan(
		&returnCategory.ID,
		&returnCategory.Name,
		&returnCategory.CreatedAt,
		&returnCategory.UpdatedAt,
	)

	if err != nil {
		if errCode := cr.db.ErrorCode(err); errCode == "23505" {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return returnCategory.ToDomain(), nil
}

// GetCategoryByID retrieves a category record from the database by id
func (cr *categoryRepository) GetCategoryByID(ctx context.Context, id uint64) (*domaincategory.Category, error) {
	query := cr.db.QueryBuilder.Select("*").
		From("categories").
		Where(sq.Eq{"id": id}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var returnCategory model.Category

	err = cr.db.QueryRow(ctx, sql, args...).Scan(
		&returnCategory.ID,
		&returnCategory.Name,
		&returnCategory.CreatedAt,
		&returnCategory.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return returnCategory.ToDomain(), nil
}

// ListCategories retrieves a list of categories from the database
func (cr *categoryRepository) ListCategories(ctx context.Context, skip, limit uint64) ([]domaincategory.Category, error) {
	var category model.Category
	var categories []domaincategory.Category

	query := cr.db.QueryBuilder.Select("*").
		From("categories").
		OrderBy("id").
		Limit(limit).
		Offset((skip - 1) * limit)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := cr.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, *category.ToDomain())
	}

	return categories, nil
}

// UpdateCategory updates a category record in the database
func (cr *categoryRepository) UpdateCategory(ctx context.Context, category *domaincategory.Category) (*domaincategory.Category, error) {
	query := cr.db.QueryBuilder.Update("categories").
		Set("name", category.Name).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": category.ID}).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var updatedCategory model.Category

	err = cr.db.QueryRow(ctx, sql, args...).Scan(
		&updatedCategory.ID,
		&updatedCategory.Name,
		&updatedCategory.CreatedAt,
		&updatedCategory.UpdatedAt,
	)
	if err != nil {
		if errCode := cr.db.ErrorCode(err); errCode == "23505" {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return updatedCategory.ToDomain(), nil
}

// DeleteCategory deletes a category record from the database by id
func (cr *categoryRepository) DeleteCategory(ctx context.Context, id uint64) error {
	query := cr.db.QueryBuilder.Delete("categories").
		Where(sq.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = cr.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
