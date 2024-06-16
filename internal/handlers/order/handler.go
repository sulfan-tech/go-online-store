package order

import (
	"net/http"
	"strconv"

	"go-online-store/internal/domain/order/service"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService service.OrderServiceImpl
}

func NewOrderHandler(orderService service.OrderServiceImpl) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) CheckoutHandler(c echo.Context) error {
	ctx := c.Request().Context()

	order, err := h.orderService.Checkout(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) TransactionPaidHandler(c echo.Context) error {
	ctx := c.Request().Context()

	orderId := c.QueryParam("order_id")

	id, _ := strconv.Atoi(orderId)

	err := h.orderService.UpdatePaymentStatus(ctx, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Payment Successfully")
}
