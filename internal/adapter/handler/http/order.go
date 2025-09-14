package http

import (
	"github.com/TienMinh25/go-hexagonal-architecture/internal/adapter/handler/http/dto"
	domainorder "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/order"
	portin "github.com/TienMinh25/go-hexagonal-architecture/internal/application/port/in"
	"github.com/gin-gonic/gin"
)

// OrderHandler represents the HTTP handler for order-related requests
type OrderHandler struct {
	svc portin.OrderService
}

// NewOrderHandler creates a new OrderHandler instance
func NewOrderHandler(svc portin.OrderService) *OrderHandler {
	return &OrderHandler{
		svc,
	}
}

// CreateOrder godoc
//
//	@Summary		Create a new order
//	@Description	Create a new order and return the order data with purchase details
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			createOrderRequest	body		dto.CreateOrderRequest	true	"Create order request"
//	@Success		200					{object}	dto.OrderResponse		"Order created"
//	@Failure		400					{object}	dto.ErrorResponse		"Validation error"
//	@Failure		404					{object}	dto.ErrorResponse		"Data not found error"
//	@Failure		409					{object}	dto.ErrorResponse		"Data conflict error"
//	@Failure		500					{object}	dto.ErrorResponse		"Internal server error"
//	@Router			/orders [post]
//	@Security		BearerAuth
func (oh *OrderHandler) CreateOrder(ctx *gin.Context) {
	var req dto.CreateOrderRequest
	var products []domainorder.OrderProduct

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	for _, product := range req.Products {
		products = append(products, domainorder.OrderProduct{
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
		})
	}

	authPayload := getAuthPayload(ctx, authorizationPayloadKey)

	order := domainorder.Order{
		UserID:       authPayload.UserID,
		PaymentID:    req.PaymentID,
		CustomerName: req.CustomerName,
		TotalPaid:    float64(req.TotalPaid),
		Products:     products,
	}

	_, err := oh.svc.CreateOrder(ctx, &order)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newOrderResponse(&order)

	handleSuccess(ctx, rsp)
}

// GetOrder godoc
//
//	@Summary		Get an order
//	@Description	Get an order by id and return the order data with purchase details
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64			true	"Order ID"
//	@Success		200	{object}	dto.OrderResponse	"Order displayed"
//	@Failure		400	{object}	dto.ErrorResponse	"Validation error"
//	@Failure		404	{object}	dto.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	dto.ErrorResponse	"Internal server error"
//	@Router			/orders/{id} [get]
//	@Security		BearerAuth
func (oh *OrderHandler) GetOrder(ctx *gin.Context) {
	var req dto.GetOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		validationError(ctx, err)
		return
	}

	order, err := oh.svc.GetOrder(ctx, req.ID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newOrderResponse(order)

	handleSuccess(ctx, rsp)
}

// ListOrders godoc
//
//	@Summary		List orders
//	@Description	List orders and return an array of order data with purchase details
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			skip	query		uint64			true	"Skip records"
//	@Param			limit	query		uint64			true	"Limit records"
//	@Success		200		{object}	dto.Meta			"Orders displayed"
//	@Failure		400		{object}	dto.ErrorResponse	"Validation error"
//	@Failure		401		{object}	dto.ErrorResponse	"Unauthorized error"
//	@Failure		500		{object}	dto.ErrorResponse	"Internal server error"
//	@Router			/orders [get]
//	@Security		BearerAuth
func (oh *OrderHandler) ListOrders(ctx *gin.Context) {
	var req dto.ListOrdersRequest
	var ordersList []dto.OrderResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	orders, err := oh.svc.ListOrders(ctx, req.Skip, req.Limit)
	if err != nil {
		handleError(ctx, err)
		return
	}

	for _, order := range orders {
		ordersList = append(ordersList, newOrderResponse(&order))
	}

	total := uint64(len(ordersList))
	meta := newMeta(total, req.Limit, req.Skip)
	rsp := toMap(meta, ordersList, "orders")

	handleSuccess(ctx, rsp)
}
