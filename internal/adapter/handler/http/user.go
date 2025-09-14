package http

import (
	domainuser "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/user"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/port"
	modelv1 "github.com/TienMinh25/go-hexagonal-architecture/pkg/model/v1"
	"github.com/gin-gonic/gin"
)

// UserHandler represents the HTTP handler for user-related requests
type UserHandler struct {
	svc port.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	create a new user account with default role "cashier"
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			registerRequest	body		modelv1.RegisterRequest	true	"Register request"
//	@Success		200				{object}	modelv1.UserResponse	"User created"
//	@Failure		400				{object}	modelv1.ErrorResponse	"Validation error"
//	@Failure		401				{object}	modelv1.ErrorResponse	"Unauthorized error"
//	@Failure		404				{object}	modelv1.ErrorResponse	"Data not found error"
//	@Failure		409				{object}	modelv1.ErrorResponse	"Data conflict error"
//	@Failure		500				{object}	modelv1.ErrorResponse	"Internal server error"
//	@Router			/users [post]
func (uh *UserHandler) Register(ctx *gin.Context) {
	var req modelv1.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	user := domainuser.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	_, err := uh.svc.Register(ctx, &user)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newUserResponse(&user)

	handleSuccess(ctx, rsp)
}

// ListUsers godoc
//
//	@Summary		List users
//	@Description	List users with pagination
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			skip	query		uint64					true	"Skip"
//	@Param			limit	query		uint64					true	"Limit"
//	@Success		200		{object}	modelv1.Meta			"Users displayed"
//	@Failure		400		{object}	modelv1.ErrorResponse	"Validation error"
//	@Failure		500		{object}	modelv1.ErrorResponse	"Internal server error"
//	@Router			/users [get]
//	@Security		BearerAuth
func (uh *UserHandler) ListUsers(ctx *gin.Context) {
	var req modelv1.ListUsersRequest
	var usersList []modelv1.UserResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	users, err := uh.svc.ListUsers(ctx, req.Skip, req.Limit)
	if err != nil {
		handleError(ctx, err)
		return
	}

	for _, user := range users {
		usersList = append(usersList, newUserResponse(&user))
	}

	total := uint64(len(usersList))
	meta := newMeta(total, req.Limit, req.Skip)
	rsp := toMap(meta, usersList, "users")

	handleSuccess(ctx, rsp)
}

// GetUser godoc
//
//	@Summary		Get a user
//	@Description	Get a user by id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64					true	"User ID"
//	@Success		200	{object}	modelv1.UserResponse	"User displayed"
//	@Failure		400	{object}	modelv1.ErrorResponse	"Validation error"
//	@Failure		404	{object}	modelv1.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	modelv1.ErrorResponse	"Internal server error"
//	@Router			/users/{id} [get]
//	@Security		BearerAuth
func (uh *UserHandler) GetUser(ctx *gin.Context) {
	var req modelv1.GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
		return
	}

	user, err := uh.svc.GetUser(ctx, req.ID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newUserResponse(user)

	handleSuccess(ctx, rsp)
}

// UpdateUser godoc
//
//	@Summary		Update a user
//	@Description	Update a user's name, email, password, or role by id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id					path		uint64						true	"User ID"
//	@Param			updateUserRequest	body		modelv1.UpdateUserRequest	true	"Update user request"
//	@Success		200					{object}	modelv1.UserResponse		"User updated"
//	@Failure		400					{object}	modelv1.ErrorResponse		"Validation error"
//	@Failure		401					{object}	modelv1.ErrorResponse		"Unauthorized error"
//	@Failure		403					{object}	modelv1.ErrorResponse		"Forbidden error"
//	@Failure		404					{object}	modelv1.ErrorResponse		"Data not found error"
//	@Failure		500					{object}	modelv1.ErrorResponse		"Internal server error"
//	@Router			/users/{id} [put]
//	@Security		BearerAuth
func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var req modelv1.UpdateUserRequest
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

	user := domainuser.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	_, err = uh.svc.UpdateUser(ctx, &user)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newUserResponse(&user)

	handleSuccess(ctx, rsp)
}

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user by id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64					true	"User ID"
//	@Success		200	{object}	modelv1.Response		"User deleted"
//	@Failure		400	{object}	modelv1.ErrorResponse	"Validation error"
//	@Failure		401	{object}	modelv1.ErrorResponse	"Unauthorized error"
//	@Failure		403	{object}	modelv1.ErrorResponse	"Forbidden error"
//	@Failure		404	{object}	modelv1.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	modelv1.ErrorResponse	"Internal server error"
//	@Router			/users/{id} [delete]
//	@Security		BearerAuth
func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	var req modelv1.DeleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
		return
	}

	err := uh.svc.DeleteUser(ctx, req.ID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}
