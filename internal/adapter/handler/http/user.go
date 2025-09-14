package http

import (
	"github.com/TienMinh25/go-hexagonal-architecture/internal/adapter/handler/http/dto"
	domainuser "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/user"
	portin "github.com/TienMinh25/go-hexagonal-architecture/internal/application/port/in"
	"github.com/gin-gonic/gin"
)

// UserHandler represents the HTTP handler for user-related requests
type UserHandler struct {
	svc portin.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(svc portin.UserService) *UserHandler {
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
//	@Param			registerRequest	body		dto.RegisterRequest	true	"Register request"
//	@Success		200				{object}	dto.UserResponse	"User created"
//	@Failure		400				{object}	dto.ErrorResponse	"Validation error"
//	@Failure		401				{object}	dto.ErrorResponse	"Unauthorized error"
//	@Failure		404				{object}	dto.ErrorResponse	"Data not found error"
//	@Failure		409				{object}	dto.ErrorResponse	"Data conflict error"
//	@Failure		500				{object}	dto.ErrorResponse	"Internal server error"
//	@Router			/users [post]
func (uh *UserHandler) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
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
//	@Param			skip	query		uint64			true	"Skip"
//	@Param			limit	query		uint64			true	"Limit"
//	@Success		200		{object}	dto.Meta			"Users displayed"
//	@Failure		400		{object}	dto.ErrorResponse	"Validation error"
//	@Failure		500		{object}	dto.ErrorResponse	"Internal server error"
//	@Router			/users [get]
//	@Security		BearerAuth
func (uh *UserHandler) ListUsers(ctx *gin.Context) {
	var req dto.ListUsersRequest
	var usersList []dto.UserResponse

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
//	@Param			id	path		uint64			true	"User ID"
//	@Success		200	{object}	dto.UserResponse	"User displayed"
//	@Failure		400	{object}	dto.ErrorResponse	"Validation error"
//	@Failure		404	{object}	dto.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	dto.ErrorResponse	"Internal server error"
//	@Router			/users/{id} [get]
//	@Security		BearerAuth
func (uh *UserHandler) GetUser(ctx *gin.Context) {
	var req dto.GetUserRequest
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
//	@Param			id					path		uint64				true	"User ID"
//	@Param			updateUserRequest	body		dto.UpdateUserRequest	true	"Update user request"
//	@Success		200					{object}	dto.UserResponse		"User updated"
//	@Failure		400					{object}	dto.ErrorResponse		"Validation error"
//	@Failure		401					{object}	dto.ErrorResponse		"Unauthorized error"
//	@Failure		403					{object}	dto.ErrorResponse		"Forbidden error"
//	@Failure		404					{object}	dto.ErrorResponse		"Data not found error"
//	@Failure		500					{object}	dto.ErrorResponse		"Internal server error"
//	@Router			/users/{id} [put]
//	@Security		BearerAuth
func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var req dto.UpdateUserRequest
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
//	@Param			id	path		uint64			true	"User ID"
//	@Success		200	{object}	dto.Response		"User deleted"
//	@Failure		400	{object}	dto.ErrorResponse	"Validation error"
//	@Failure		401	{object}	dto.ErrorResponse	"Unauthorized error"
//	@Failure		403	{object}	dto.ErrorResponse	"Forbidden error"
//	@Failure		404	{object}	dto.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	dto.ErrorResponse	"Internal server error"
//	@Router			/users/{id} [delete]
//	@Security		BearerAuth
func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	var req dto.DeleteUserRequest
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
