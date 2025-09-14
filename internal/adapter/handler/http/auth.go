package http

import (
	portin "github.com/TienMinh25/go-hexagonal-architecture/internal/application/port/in"
	modelv1 "github.com/TienMinh25/go-hexagonal-architecture/pkg/model/v1"
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
//	@Param			request	body		modelv1.LoginRequest	true	"Login request body"
//	@Success		200		{object}	modelv1.AuthResponse	"Succesfully logged in"
//	@Failure		400		{object}	modelv1.ErrorResponse	"Validation error"
//	@Failure		401		{object}	modelv1.ErrorResponse	"Unauthorized error"
//	@Failure		500		{object}	modelv1.ErrorResponse	"Internal server error"
//	@Router			/users/login [post]
func (ah *AuthHandler) Login(ctx *gin.Context) {
	var req modelv1.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	token, err := ah.svc.Login(ctx, req.Email, req.Password)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newAuthResponse(token)

	handleSuccess(ctx, rsp)
}
