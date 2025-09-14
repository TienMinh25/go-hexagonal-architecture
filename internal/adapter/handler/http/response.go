package http

import (
	"errors"
	"net/http"

	"github.com/TienMinh25/go-hexagonal-architecture/internal/adapter/handler/http/dto"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain"
	domaincategory "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/category"
	domainorder "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/order"
	domainpayment "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/payment"
	domainproduct "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/product"
	domainuser "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// newResponse is a helper function to create a response body
func newResponse(success bool, message string, data any) dto.Response {
	return dto.Response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

// newMeta is a helper function to create metadata for a paginated response
func newMeta(total, limit, skip uint64) dto.Meta {
	return dto.Meta{
		Total: total,
		Limit: limit,
		Skip:  skip,
	}
}

// newUserResponse is a helper function to create a response body for handling user data
func newUserResponse(user *domainuser.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// newPaymentResponse is a helper function to create a response body for handling payment data
func newPaymentResponse(payment *domainpayment.Payment) dto.PaymentResponse {
	return dto.PaymentResponse{
		ID:   payment.ID,
		Name: payment.Name,
		Type: payment.Type,
		Logo: payment.Logo,
	}
}

// newCategoryResponse is a helper function to create a response body for handling category data
func newCategoryResponse(category *domaincategory.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

// newProductResponse is a helper function to create a response body for handling product data
func newProductResponse(product *domainproduct.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:        product.ID,
		SKU:       product.SKU.String(),
		Name:      product.Name,
		Stock:     product.Stock,
		Price:     product.Price,
		Image:     product.Image,
		Category:  newCategoryResponse(product.Category),
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

// newOrderResponse is a helper function to create a response body for handling order data
func newOrderResponse(order *domainorder.Order) dto.OrderResponse {
	return dto.OrderResponse{
		ID:           order.ID,
		UserID:       order.UserID,
		PaymentID:    order.PaymentID,
		CustomerName: order.CustomerName,
		TotalPrice:   order.TotalPrice,
		TotalPaid:    order.TotalPaid,
		TotalReturn:  order.TotalReturn,
		ReceiptCode:  order.ReceiptCode.String(),
		Products:     newOrderProductResponse(order.Products),
		PaymentType:  newPaymentResponse(order.Payment),
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}
}

// newOrderProductResponse is a helper function to create a response body for handling order product data
func newOrderProductResponse(orderProduct []domainorder.OrderProduct) []dto.OrderProductResponse {
	var orderProductResponses []dto.OrderProductResponse

	for _, orderProduct := range orderProduct {
		orderProductResponses = append(orderProductResponses, dto.OrderProductResponse{
			ID:               orderProduct.ID,
			OrderID:          orderProduct.OrderID,
			ProductID:        orderProduct.ProductID,
			Quantity:         orderProduct.Quantity,
			Price:            orderProduct.Product.Price,
			TotalNormalPrice: orderProduct.TotalPrice,
			TotalFinalPrice:  orderProduct.TotalPrice,
			Product:          newProductResponse(orderProduct.Product),
			CreatedAt:        orderProduct.CreatedAt,
			UpdatedAt:        orderProduct.UpdatedAt,
		})
	}

	return orderProductResponses
}

// errorStatusMap is a map of defined error messages and their corresponding http status codes
var errorStatusMap = map[error]int{
	domain.ErrInternal:                   http.StatusInternalServerError,
	domain.ErrDataNotFound:               http.StatusNotFound,
	domain.ErrConflictingData:            http.StatusConflict,
	domain.ErrInvalidCredentials:         http.StatusUnauthorized,
	domain.ErrUnauthorized:               http.StatusUnauthorized,
	domain.ErrEmptyAuthorizationHeader:   http.StatusUnauthorized,
	domain.ErrInvalidAuthorizationHeader: http.StatusUnauthorized,
	domain.ErrInvalidAuthorizationType:   http.StatusUnauthorized,
	domain.ErrInvalidToken:               http.StatusUnauthorized,
	domain.ErrExpiredToken:               http.StatusUnauthorized,
	domain.ErrForbidden:                  http.StatusForbidden,
	domain.ErrNoUpdatedData:              http.StatusBadRequest,
	domain.ErrInsufficientStock:          http.StatusBadRequest,
	domain.ErrInsufficientPayment:        http.StatusBadRequest,
}

// validationError sends an error response for some specific request validation error
func validationError(ctx *gin.Context, err error) {
	errMsgs := parseError(err)
	errRsp := newErrorResponse(errMsgs)
	ctx.JSON(http.StatusBadRequest, errRsp)
}

// handleError determines the status code of an error and returns a JSON response with the error message and status code
func handleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := newErrorResponse(errMsg)
	ctx.JSON(statusCode, errRsp)
}

// handleAbort sends an error response and aborts the request with the specified status code and error message
func handleAbort(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := newErrorResponse(errMsg)
	ctx.AbortWithStatusJSON(statusCode, errRsp)
}

// parseError parses error messages from the error object and returns a slice of error messages
func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

// NewErrorResponse is a helper function to create an error response body
func newErrorResponse(errMsgs []string) dto.ErrorResponse {
	return dto.ErrorResponse{
		Success:  false,
		Messages: errMsgs,
	}
}

// handleSuccess sends a success response with the specified status code and optional data
func handleSuccess(ctx *gin.Context, data any) {
	rsp := newResponse(true, "Success", data)
	ctx.JSON(http.StatusOK, rsp)
}
