package http

import (
	domainpayment "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/payment"
	"github.com/TienMinh25/go-hexagonal-architecture/internal/application/port"
	modelv1 "github.com/TienMinh25/go-hexagonal-architecture/pkg/model/v1"
	"github.com/gin-gonic/gin"
)

// PaymentHandler represents the HTTP handler for payment-related requests
type PaymentHandler struct {
	svc port.PaymentService
}

// NewPaymentHandler creates a new PaymentHandler instance
func NewPaymentHandler(svc port.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		svc,
	}
}

// CreatePayment godoc
//
//	@Summary		Create a new payment
//	@Description	create a new payment with name, type, and logo
//	@Tags			Payments
//	@Accept			json
//	@Produce		json
//	@Param			createPaymentRequest	body		modelv1.CreatePaymentRequest	true	"Create payment request"
//	@Success		200						{object}	modelv1.PaymentResponse			"Payment created"
//	@Failure		400						{object}	modelv1.ErrorResponse			"Validation error"
//	@Failure		401						{object}	modelv1.ErrorResponse			"Unauthorized error"
//	@Failure		403						{object}	modelv1.ErrorResponse			"Forbidden error"
//	@Failure		404						{object}	modelv1.ErrorResponse			"Data not found error"
//	@Failure		409						{object}	modelv1.ErrorResponse			"Data conflict error"
//	@Failure		500						{object}	modelv1.ErrorResponse			"Internal server error"
//	@Router			/payments [post]
//	@Security		BearerAuth
func (ph *PaymentHandler) CreatePayment(ctx *gin.Context) {
	var req modelv1.CreatePaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	payment := domainpayment.Payment{
		Name: req.Name,
		Type: req.Type,
		Logo: req.Logo,
	}

	_, err := ph.svc.CreatePayment(ctx, &payment)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newPaymentResponse(&payment)

	handleSuccess(ctx, rsp)
}

// GetPayment godoc
//
//	@Summary		Get a payment
//	@Description	get a payment by id
//	@Tags			Payments
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int						true	"Payment ID"
//	@Success		200	{object}	modelv1.PaymentResponse	"Payment retrieved"
//	@Failure		400	{object}	modelv1.ErrorResponse	"Validation error"
//	@Failure		404	{object}	modelv1.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	modelv1.ErrorResponse	"Internal server error"
//	@Router			/payments/{id} [get]
//	@Security		BearerAuth
func (ph *PaymentHandler) GetPayment(ctx *gin.Context) {
	var req modelv1.GetPaymentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
		return
	}

	payment, err := ph.svc.GetPayment(ctx, req.ID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newPaymentResponse(payment)

	handleSuccess(ctx, rsp)
}

// ListPayments godoc
//
//	@Summary		List payments
//	@Description	List payments with pagination
//	@Tags			Payments
//	@Accept			json
//	@Produce		json
//	@Param			skip	query		uint64					true	"Skip"
//	@Param			limit	query		uint64					true	"Limit"
//	@Success		200		{object}	modelv1.Meta			"Payments displayed"
//	@Failure		400		{object}	modelv1.ErrorResponse	"Validation error"
//	@Failure		500		{object}	modelv1.ErrorResponse	"Internal server error"
//	@Router			/payments [get]
//	@Security		BearerAuth
func (ph *PaymentHandler) ListPayments(ctx *gin.Context) {
	var req modelv1.ListPaymentsRequest
	var paymentsList []modelv1.PaymentResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	payments, err := ph.svc.ListPayments(ctx, req.Skip, req.Limit)
	if err != nil {
		handleError(ctx, err)
		return
	}

	for _, payment := range payments {
		paymentsList = append(paymentsList, newPaymentResponse(&payment))
	}

	total := uint64(len(paymentsList))
	meta := newMeta(total, req.Limit, req.Skip)
	rsp := toMap(meta, paymentsList, "payments")

	handleSuccess(ctx, rsp)
}

// UpdatePayment godoc
//
//	@Summary		Update a payment
//	@Description	update a payment's name, type, or logo by id
//	@Tags			Payments
//	@Accept			json
//	@Produce		json
//	@Param			id						path		int								true	"Payment ID"
//	@Param			updatePaymentRequest	body		modelv1.UpdatePaymentRequest	true	"Update payment request"
//	@Success		200						{object}	modelv1.PaymentResponse			"Payment updated"
//	@Failure		400						{object}	modelv1.ErrorResponse			"Validation error"
//	@Failure		401						{object}	modelv1.ErrorResponse			"Unauthorized error"
//	@Failure		403						{object}	modelv1.ErrorResponse			"Forbidden error"
//	@Failure		404						{object}	modelv1.ErrorResponse			"Data not found error"
//	@Failure		409						{object}	modelv1.ErrorResponse			"Data conflict error"
//	@Failure		500						{object}	modelv1.ErrorResponse			"Internal server error"
//	@Router			/payments/{id} [put]
//	@Security		BearerAuth
func (ph *PaymentHandler) UpdatePayment(ctx *gin.Context) {
	var req modelv1.UpdatePaymentRequest
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

	payment := domainpayment.Payment{
		ID:   id,
		Name: req.Name,
		Type: req.Type,
		Logo: req.Logo,
	}

	_, err = ph.svc.UpdatePayment(ctx, &payment)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newPaymentResponse(&payment)

	handleSuccess(ctx, rsp)
}

// DeletePayment godoc
//
//	@Summary		Delete a payment
//	@Description	Delete a payment by id
//	@Tags			Payments
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64					true	"Payment ID"
//	@Success		200	{object}	modelv1.Response		"Payment deleted"
//	@Failure		400	{object}	modelv1.ErrorResponse	"Validation error"
//	@Failure		401	{object}	modelv1.ErrorResponse	"Unauthorized error"
//	@Failure		403	{object}	modelv1.ErrorResponse	"Forbidden error"
//	@Failure		404	{object}	modelv1.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	modelv1.ErrorResponse	"Internal server error"
//	@Router			/payments/{id} [delete]
//	@Security		BearerAuth
func (ph *PaymentHandler) DeletePayment(ctx *gin.Context) {
	var req modelv1.DeletePaymentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
		return
	}

	err := ph.svc.DeletePayment(ctx, req.ID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}
