package http

import (
	domaincategory "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/category"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/port"
	modelv1 "github.com/TienMinh25/go-hexagonal-architecture/pkg/model/v1"
	"github.com/gin-gonic/gin"
)

// CategoryHandler represents the HTTP handler for category-related requests
type CategoryHandler struct {
	svc port.CategoryService
}

// NewCategoryHandler creates a new CategoryHandler instance
func NewCategoryHandler(svc port.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		svc,
	}
}

// CreateCategory godoc
//
//	@Summary		Create a new category
//	@Description	create a new category with name
//	@Tags			Categories
//	@Accept			json
//	@Produce		json
//	@Param			createCategoryRequest	body		modelv1.CreateCategoryRequest	true	"Create category request"
//	@Success		200						{object}	modelv1.CategoryResponse		"Category created"
//	@Failure		400						{object}	modelv1.ErrorResponse			"Validation error"
//	@Failure		401						{object}	modelv1.ErrorResponse			"Unauthorized error"
//	@Failure		403						{object}	modelv1.ErrorResponse			"Forbidden error"
//	@Failure		404						{object}	modelv1.ErrorResponse			"Data not found error"
//	@Failure		409						{object}	modelv1.ErrorResponse			"Data conflict error"
//	@Failure		500						{object}	modelv1.ErrorResponse			"Internal server error"
//	@Router			/categories [post]
//	@Security		BearerAuth
func (ch *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var req modelv1.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	category := domaincategory.Category{
		Name: req.Name,
	}

	_, err := ch.svc.CreateCategory(ctx, &category)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newCategoryResponse(&category)

	handleSuccess(ctx, rsp)
}

// GetCategory godoc
//
//	@Summary		Get a category
//	@Description	get a category by id
//	@Tags			Categories
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64						true	"Category ID"
//	@Success		200	{object}	modelv1.CategoryResponse	"Category retrieved"
//	@Failure		400	{object}	modelv1.ErrorResponse		"Validation error"
//	@Failure		404	{object}	modelv1.ErrorResponse		"Data not found error"
//	@Failure		500	{object}	modelv1.ErrorResponse		"Internal server error"
//	@Router			/categories/{id} [get]
//	@Security		BearerAuth
func (ch *CategoryHandler) GetCategory(ctx *gin.Context) {
	var req modelv1.GetCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
		return
	}

	category, err := ch.svc.GetCategory(ctx, req.ID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newCategoryResponse(category)

	handleSuccess(ctx, rsp)
}

// ListCategories godoc
//
//	@Summary		List categories
//	@Description	List categories with pagination
//	@Tags			Categories
//	@Accept			json
//	@Produce		json
//	@Param			skip	query		uint64					true	"Skip"
//	@Param			limit	query		uint64					true	"Limit"
//	@Success		200		{object}	modelv1.Meta			"Categories displayed"
//	@Failure		400		{object}	modelv1.ErrorResponse	"Validation error"
//	@Failure		500		{object}	modelv1.ErrorResponse	"Internal server error"
//	@Router			/categories [get]
//	@Security		BearerAuth
func (ch *CategoryHandler) ListCategories(ctx *gin.Context) {
	var req modelv1.ListCategoriesRequest
	var categoriesList []modelv1.CategoryResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	categories, err := ch.svc.ListCategories(ctx, req.Skip, req.Limit)
	if err != nil {
		handleError(ctx, err)
		return
	}

	for _, category := range categories {
		categoriesList = append(categoriesList, newCategoryResponse(&category))
	}

	total := uint64(len(categoriesList))
	meta := newMeta(total, req.Limit, req.Skip)
	rsp := toMap(meta, categoriesList, "categories")

	handleSuccess(ctx, rsp)
}

// UpdateCategory godoc
//
//	@Summary		Update a category
//	@Description	update a category's name by id
//	@Tags			Categories
//	@Accept			json
//	@Produce		json
//	@Param			id						path		uint64							true	"Category ID"
//	@Param			updateCategoryRequest	body		modelv1.UpdateCategoryRequest	true	"Update category request"
//	@Success		200						{object}	modelv1.CategoryResponse		"Category updated"
//	@Failure		400						{object}	modelv1.ErrorResponse			"Validation error"
//	@Failure		401						{object}	modelv1.ErrorResponse			"Unauthorized error"
//	@Failure		403						{object}	modelv1.ErrorResponse			"Forbidden error"
//	@Failure		404						{object}	modelv1.ErrorResponse			"Data not found error"
//	@Failure		409						{object}	modelv1.ErrorResponse			"Data conflict error"
//	@Failure		500						{object}	modelv1.ErrorResponse			"Internal server error"
//	@Router			/categories/{id} [put]
//	@Security		BearerAuth
func (ch *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	var req modelv1.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	idStr := ctx.Param("id")
	id, err := stringToUint64(idStr)
	if err != nil {
		validationError(ctx, err)
		return
	}

	category := domaincategory.Category{
		ID:   id,
		Name: req.Name,
	}

	_, err = ch.svc.UpdateCategory(ctx, &category)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newCategoryResponse(&category)

	handleSuccess(ctx, rsp)
}

// DeleteCategory godoc
//
//	@Summary		Delete a category
//	@Description	Delete a category by id
//	@Tags			Categories
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64					true	"Category ID"
//	@Success		200	{object}	modelv1.Response		"Category deleted"
//	@Failure		400	{object}	modelv1.ErrorResponse	"Validation error"
//	@Failure		401	{object}	modelv1.ErrorResponse	"Unauthorized error"
//	@Failure		403	{object}	modelv1.ErrorResponse	"Forbidden error"
//	@Failure		404	{object}	modelv1.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	modelv1.ErrorResponse	"Internal server error"
//	@Router			/categories/{id} [delete]
//	@Security		BearerAuth
func (ch *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	var req modelv1.DeleteCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
		return
	}

	err := ch.svc.DeleteCategory(ctx, req.ID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}
