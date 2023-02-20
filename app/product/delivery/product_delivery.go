package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
)

type roductDelivery struct {
	productHandler domain.IProductHandler
}

func NewProductDelivery(rg *gin.RouterGroup, handler domain.IProductHandler) {
	delivery := &roductDelivery{productHandler: handler}
	router := rg.Group("/products")

	rg.GET("/categories", delivery.GetProductCategoryList)
	router.GET("/", delivery.GetProductList)
	router.GET("/:product_id", delivery.GetProductDetail)
}

func (h *roductDelivery) GetProductList(c *gin.Context) {
	size, _ := strconv.Atoi(c.Query("size"))
	page, _ := strconv.Atoi(c.Query("page"))
	category, _ := strconv.Atoi(c.Query("category_id"))

	// Get list of products
	products, meta, err := h.productHandler.GetProductList(c, size, page, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.HttpResponseError(http.StatusInternalServerError, err.Error()))
		return
	}

	// Handle empty product
	if len(products) == 0 {
		c.JSON(http.StatusNotFound, utils.HttpResponseError(http.StatusNotFound, "no product found"))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponseWithMetaData(http.StatusOK, products, meta))
}

func (h *roductDelivery) GetProductDetail(c *gin.Context) {
	productID := c.Param("product_id")

	product, err := h.productHandler.GetProductDetail(c, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, product))
}

func (h *roductDelivery) GetProductCategoryList(c *gin.Context) {
	categories, err := h.productHandler.GetProductCategoryList(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.HttpResponseError(http.StatusInternalServerError, err.Error()))
		return
	}

	// Handle empty categories
	if len(categories) == 0 {
		c.JSON(http.StatusNotFound, utils.HttpResponseError(http.StatusNotFound, "no product categories found"))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, categories))
}
