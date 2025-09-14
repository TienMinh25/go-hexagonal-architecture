package http

import (
	domainproduct "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/product"
	portin "github.com/TienMinh25/go-hexagonal-architecture/internal/application/port/in"
	modelv1 "github.com/TienMinh25/go-hexagonal-architecture/pkg/model/v1"
	"github.com/gin-gonic/gin"
)

// ProductHandler represents the HTTP handler for product-related requests
type ProductHandler struct {
	svc portin.ProductService
}

// NewProductHandler creates a new ProductHandler instance
func NewProductHandler(svc portin.ProductService) *ProductHandler {
	return &ProductHandler{
		svc,
	}
}

// CreateProduct godoc
//
//	@Summary		Create a new product
//	@Description	create a new product with name, image, price, and stock
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			createProductRequest	body		modelv1.CreateProductRequest	true	"Create product request"
//	@Success		200						{object}	modelv1.ProductResponse			"Product created"
//	@Failure		400						{object}	modelv1.ErrorResponse			"Validation error"
//	@Failure		401						{object}	modelv1.ErrorResponse			"Unauthorized error"
//	@Failure		403						{object}	modelv1.ErrorResponse			"Forbidden error"
//	@Failure		404						{object}	modelv1.ErrorResponse			"Data not found error"
//	@Failure		409						{object}	modelv1.ErrorResponse			"Data conflict error"
//	@Failure		500						{object}	modelv1.ErrorResponse			"Internal server error"
//	@Router			/products [post]
//	@Security		BearerAuth
func (ph *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req modelv1.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	product := domainproduct.Product{
		CategoryID: req.CategoryID,
		Name:       req.Name,
		Image:      req.Image,
		Price:      req.Price,
		Stock:      req.Stock,
	}

	_, err := ph.svc.CreateProduct(ctx, &product)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newProductResponse(&product)

	handleSuccess(ctx, rsp)
}

// GetProduct godoc
//
//	@Summary		Get a product
//	@Description	get a product by id with its category
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64					true	"Product ID"
//	@Success		200	{object}	modelv1.ProductResponse	"Product retrieved"
//	@Failure		400	{object}	modelv1.ErrorResponse	"Validation error"
//	@Failure		404	{object}	modelv1.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	modelv1.ErrorResponse	"Internal server error"
//	@Router			/products/{id} [get]
//	@Security		BearerAuth
func (ph *ProductHandler) GetProduct(ctx *gin.Context) {
	var req modelv1.GetProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
		return
	}

	product, err := ph.svc.GetProduct(ctx, req.ID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newProductResponse(product)

	handleSuccess(ctx, rsp)
}

// ListProducts godoc
//
//	@Summary		List products
//	@Description	List products with pagination
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			category_id	query		uint64					false	"Category ID"
//	@Param			q			query		string					false	"Query"
//	@Param			skip		query		uint64					true	"Skip"
//	@Param			limit		query		uint64					true	"Limit"
//	@Success		200			{object}	modelv1.Meta			"Products retrieved"
//	@Failure		400			{object}	modelv1.ErrorResponse	"Validation error"
//	@Failure		500			{object}	modelv1.ErrorResponse	"Internal server error"
//	@Router			/products [get]
//	@Security		BearerAuth
func (ph *ProductHandler) ListProducts(ctx *gin.Context) {
	var req modelv1.ListProductsRequest
	var productsList []modelv1.ProductResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	products, err := ph.svc.ListProducts(ctx, req.Query, req.CategoryID, req.Skip, req.Limit)
	if err != nil {
		handleError(ctx, err)
		return
	}

	for _, product := range products {
		productsList = append(productsList, newProductResponse(&product))
	}

	total := uint64(len(productsList))
	meta := newMeta(total, req.Limit, req.Skip)
	rsp := toMap(meta, productsList, "products")

	handleSuccess(ctx, rsp)
}

// UpdateProduct godoc
//
//	@Summary		Update a product
//	@Description	update a product's name, image, price, or stock by id
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id						path		uint64							true	"Product ID"
//	@Param			updateProductRequest	body		modelv1.UpdateProductRequest	true	"Update product request"
//	@Success		200						{object}	modelv1.ProductResponse			"Product updated"
//	@Failure		400						{object}	modelv1.ErrorResponse			"Validation error"
//	@Failure		401						{object}	modelv1.ErrorResponse			"Unauthorized error"
//	@Failure		403						{object}	modelv1.ErrorResponse			"Forbidden error"
//	@Failure		404						{object}	modelv1.ErrorResponse			"Data not found error"
//	@Failure		409						{object}	modelv1.ErrorResponse			"Data conflict error"
//	@Failure		500						{object}	modelv1.ErrorResponse			"Internal server error"
//	@Router			/products/{id} [put]
//	@Security		BearerAuth
func (ph *ProductHandler) UpdateProduct(ctx *gin.Context) {
	var req modelv1.UpdateProductRequest
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

	product := domainproduct.Product{
		ID:         id,
		CategoryID: req.CategoryID,
		Name:       req.Name,
		Image:      req.Image,
		Price:      req.Price,
		Stock:      req.Stock,
	}

	_, err = ph.svc.UpdateProduct(ctx, &product)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newProductResponse(&product)

	handleSuccess(ctx, rsp)
}

// DeleteProduct godoc
//
//	@Summary		Delete a product
//	@Description	Delete a product by id
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64					true	"Product ID"
//	@Success		200	{object}	modelv1.Response		"Product deleted"
//	@Failure		400	{object}	modelv1.ErrorResponse	"Validation error"
//	@Failure		401	{object}	modelv1.ErrorResponse	"Unauthorized error"
//	@Failure		403	{object}	modelv1.ErrorResponse	"Forbidden error"
//	@Failure		404	{object}	modelv1.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	modelv1.ErrorResponse	"Internal server error"
//	@Router			/products/{id} [delete]
//	@Security		BearerAuth
func (ph *ProductHandler) DeleteProduct(ctx *gin.Context) {
	var req modelv1.DeleteProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
		return
	}

	err := ph.svc.DeleteProduct(ctx, req.ID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}
