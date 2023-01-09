package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
)

type ProductHandler struct {
	productService domain.IProductService
}

func NewProductHandler(rg *gin.RouterGroup, service domain.IProductService) {
	handler := &ProductHandler{productService: service}
	router := rg.Group("/products")

	router.GET("/", handler.GetProductList)
}

func (h *ProductHandler) GetProductList(c *gin.Context) {
	pageParam := c.Query("page")
	category := c.Query("category")

	// Validate page param
	page, err := strconv.Atoi(pageParam)
	if err != nil && pageParam != "" {
		errMsg := map[string]string{
			"message": "invalid page param type, expected int",
		}
		c.JSON(http.StatusBadRequest, utils.HttpResponse(http.StatusBadRequest, errMsg, nil))
		return
	}
	if pageParam == "" {
		page = 1
	}

	// Get list of products
	products, err := h.productService.GetProductList(c, page, category)
	if err != nil {
		errMsg := map[string]string{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, utils.HttpResponse(http.StatusInternalServerError, errMsg, nil))
		return
	}

	// Handle empty product
	if len(products) == 0 {
		errMsg := map[string]string{
			"message": "no product found",
		}
		c.JSON(http.StatusNotFound, utils.HttpResponse(http.StatusNotFound, errMsg, []domain.ProductDTO{}))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, nil, products))
}
