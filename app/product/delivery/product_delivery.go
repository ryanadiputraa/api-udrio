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
		c.JSON(http.StatusNotFound, utils.HttpResponse(http.StatusNotFound, errMsg, []domain.Product{}))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponseWithMetaData(http.StatusOK, nil, products, meta))
}

func (h *roductDelivery) GetProductDetail(c *gin.Context) {
	productID := c.Param("product_id")

	product, err := h.productHandler.GetProductDetail(c, productID)
	if err != nil {
		errMsg := map[string]string{
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, utils.HttpResponse(http.StatusBadRequest, errMsg, nil))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, nil, product))
}

func (h *roductDelivery) GetProductCategoryList(c *gin.Context) {
	categories, err := h.productHandler.GetProductCategoryList(c)
	if err != nil {
		errMsg := map[string]string{
			"message": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, utils.HttpResponse(http.StatusInternalServerError, errMsg, nil))
		return
	}

	// Handle empty categories
	if len(categories) == 0 {
		errMsg := map[string]string{
			"message": "no product categories found",
		}
		c.JSON(http.StatusNotFound, utils.HttpResponse(http.StatusNotFound, errMsg, []domain.ProductCategory{}))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, nil, categories))
}
