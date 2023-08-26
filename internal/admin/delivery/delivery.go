package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/config"
	"github.com/ryanadiputraa/api-udrio/domain"
)

type delivery struct {
	config         config.Config
	usecase        domain.AdminUsecase
	productUsecase domain.ProductUsecase
}

func NewAdminDelivery(rg *gin.RouterGroup, config config.Config, usecase domain.AdminUsecase, productUsecase domain.ProductUsecase) {
	delivery := delivery{config: config, usecase: usecase, productUsecase: productUsecase}
	rg.GET("/", delivery.parseSessionToken(), delivery.MainPanel)
	rg.GET("/login", delivery.Login)
	rg.POST("/signin", delivery.SignIn)
	rg.GET("/products", delivery.parseSessionToken(), delivery.Products)
	rg.POST("/products", delivery.parseSessionToken(), delivery.UploadProducts)
	rg.GET("/products/:id", delivery.parseSessionToken(), delivery.ProductDetail)
	rg.POST("/products/:id", delivery.parseSessionToken(), delivery.UpdateProduct)
}

func (d *delivery) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func (d *delivery) SignIn(c *gin.Context) {
	c.Request.ParseForm()
	sessionToken, expiresAt, err := d.usecase.SignIn(c, c.PostForm("username"), c.PostForm("password"))
	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})
	c.Redirect(http.StatusFound, "/admin/products")
}

func (d *delivery) Products(c *gin.Context) {
	path, _ := d.usecase.GetFilePath(c, "products")
	c.HTML(http.StatusOK, "products.html", gin.H{
		"filepath": path.FilePath,
	})
}

func (d *delivery) ProductDetail(c *gin.Context) {
	product, _ := d.productUsecase.GetProductDetail(c, c.Param("id"))
	stock := "Ada"
	if !product.IsAvailable {
		stock = "Kosong"
	}

	c.HTML(http.StatusOK, "product-detail.html", gin.H{
		"id":           product.ID,
		"product_name": product.ProductName,
		"category_id":  product.ProductCategoryID,
		"price":        product.Price,
		"stock":        stock,
		"description":  product.Description,
		"min_order":    product.MinOrder,
		"images":       product.ProductImages,
	})
}

func (d *delivery) MainPanel(c *gin.Context) {
	c.Redirect(http.StatusFound, "/admin/products")
}

func (d *delivery) UploadProducts(c *gin.Context) {
	path, _ := d.usecase.GetFilePath(c, "products")

	file, err := c.FormFile("file")
	if err != nil {
		c.HTML(http.StatusOK, "products.html", gin.H{
			"error":    err.Error(),
			"filepath": path.FilePath,
		})
		return
	}

	filePath := "assets/uploads/" + file.Filename
	assetsPath := domain.AssetsPath{
		Key:      "products",
		FilePath: filePath,
	}
	if err = d.usecase.SaveFilePath(c, assetsPath); err != nil {
		c.HTML(http.StatusOK, "products.html", gin.H{
			"error":    err.Error(),
			"filepath": path.FilePath,
		})
		return
	}

	if err = c.SaveUploadedFile(file, filePath); err != nil {
		assetsPath.FilePath = ""
		d.usecase.SaveFilePath(c, assetsPath)
		c.HTML(http.StatusOK, "products.html", gin.H{
			"error":    err.Error(),
			"filepath": path.FilePath,
		})
		return
	}

	if err = d.usecase.BulkInsertProducts(c); err != nil {
		c.HTML(http.StatusOK, "products.html", gin.H{
			"error":    err.Error(),
			"filepath": path.FilePath,
		})
		return
	}

	c.HTML(http.StatusOK, "products.html", gin.H{
		"message":  "Data berhasil diperbarui",
		"filepath": filePath,
	})
}

func (d *delivery) UpdateProduct(c *gin.Context) {
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		d.ProductDetail(c)
		return
	}
	defer file.Close()

	productID := c.Request.FormValue("product-id")
	d.productUsecase.UploadProductImage(c, productID, file)
	d.ProductDetail(c)
}

func (d *delivery) parseSessionToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken, err := c.Cookie("session_token")
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/admin/login")
		}

		_, err = d.usecase.GetSession(c, sessionToken)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/admin/login")
		}
		c.Next()
	}
}
