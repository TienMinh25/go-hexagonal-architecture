package http

import (
	"github.com/TienMinh25/go-hexagonal-architecture/internal/adapter/handler/http/dto"
	portin "github.com/TienMinh25/go-hexagonal-architecture/internal/application/port/in"
	"github.com/gin-gonic/gin"
)

// AuthHandler represents the HTTP handler for authentication-related requests
type AuthHandler struct {
	svc portin.AuthService
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler(svc portin.AuthService) *AuthHandler {
	return &AuthHandler{
		svc,
	}
}

// Login godoc
//
//	@Summary		Login and get an access token
//	@Description	Logs in a registered user and returns an access token if the credentials are valid.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LoginRequest	true	"Login request body"
//	@Success		200		{object}	dto.AuthResponse	"Succesfully logged in"
//	@Failure		400		{object}	dto.ErrorResponse	"Validation error"
//	@Failure		401		{object}	dto.ErrorResponse	"Unauthorized error"
//	@Failure		500		{object}	dto.ErrorResponse	"Internal server error"
//	@Router			/users/login [post]
func (ah *AuthHandler) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	token, err := ah.svc.Login(ctx, req.Email, req.Password)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := dto.NewAuthResponse(token)

	handleSuccess(ctx, rsp)
}
