package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
)

type delivery struct {
	usecase domain.ProductUsecase
}

func NewProductDelivery(rg *gin.RouterGroup, usecase domain.ProductUsecase) {
	delivery := &delivery{usecase: usecase}
	router := rg.Group("/products")

	rg.GET("/categories", delivery.GetProductCategoryList)
	router.GET("/", delivery.GetProductList)
	router.GET("/:product_id", delivery.GetProductDetail)
}

func (d *delivery) GetProductList(c *gin.Context) {
	size, _ := strconv.Atoi(c.Query("size"))
	page, _ := strconv.Atoi(c.Query("page"))
	category, _ := strconv.Atoi(c.Query("category_id"))
	query := c.Query("query")

	// Get list of products
	products, meta, err := d.usecase.GetProductList(c, size, page, category, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.HttpResponseError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponseWithMetaData(http.StatusOK, products, meta))
}

func (d *delivery) GetProductDetail(c *gin.Context) {
	productID := c.Param("product_id")

	product, err := d.usecase.GetProductDetail(c, productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HttpResponseError(http.StatusBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, product))
}

func (d *delivery) GetProductCategoryList(c *gin.Context) {
	categories, err := d.usecase.GetProductCategoryList(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.HttpResponseError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.HttpResponse(http.StatusOK, categories))
}
