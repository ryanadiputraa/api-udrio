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

func NewProductHandler(rg *gin.RouterGroup) {
	handler := &ProductHandler{}
	router := rg.Group("/products")

	router.GET("/", handler.GetProductList)
}

func (h *ProductHandler) GetProductList(c *gin.Context) {
	pageParam := c.Query("page")
	_ = c.Query("category")

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

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, nil, page))
}
