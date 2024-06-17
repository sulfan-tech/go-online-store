package product

import (
	"net/http"

	"go-online-store/internal/domain/product/service"
	"go-online-store/pkg/errors"
	_ "go-online-store/server/cmd/docs"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService service.ProductServiceImpl
}

func NewProductHandler(productService service.ProductServiceImpl) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// @Summary Get products by category
// @Tags Product
// @Description Retrieve a list of products filtered by category
// @Produce json
// @Param category query string false "Category filter"
// @Success 200 {object} []Product
// @Failure 400 {object} ErrorResponse
// @Router /v1/products [get]

// GetProductsByCategoryHandler handles the request to fetch products by category
func (h *ProductHandler) GetProductsByCategoryHandler(c echo.Context) error {

	ctx := c.Request().Context()

	category := c.QueryParam("category")
	if category == "" {
		return errors.HTTPErrorHandler(errors.ErrBadRequest)
	}

	products, err := h.productService.GetProductListByCategory(ctx, category)
	if err != nil {
		return errors.HTTPErrorHandler(err)
	}

	return c.JSON(http.StatusOK, products)
}
