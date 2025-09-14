package dto

// CreateCategoryRequest represents a request body for creating a new category
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required" example:"Foods"`
}

// categoryResponse represents a category response body
type CategoryResponse struct {
	ID   uint64 `json:"id" example:"1"`
	Name string `json:"name" example:"Foods"`
}

// GetCategoryRequest represents a request body for retrieving a category
type GetCategoryRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}

// ListCategoriesRequest represents a request body for listing categories
type ListCategoriesRequest struct {
	Skip  uint64 `form:"skip" binding:"required,min=0" example:"0"`
	Limit uint64 `form:"limit" binding:"required,min=5" example:"5"`
}

// UpdateCategoryRequest represents a request body for updating a category
type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"omitempty,required" example:"Beverages"`
}

// DeleteCategoryRequest represents a request body for deleting a category
type DeleteCategoryRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}
