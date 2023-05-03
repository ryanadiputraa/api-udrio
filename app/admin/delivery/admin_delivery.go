package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/domain"
)

type adminDelivery struct {
	handler domain.IAdminHandler
}

func NewAdminDelivery(rg *gin.RouterGroup, handler domain.IAdminHandler) {
	delivery := adminDelivery{handler: handler}
	rg.GET("/", delivery.parseSessionToken(), delivery.MainPanel)
	rg.GET("/login", delivery.Login)
	rg.POST("/signin", delivery.SignIn)
	rg.GET("/products", delivery.parseSessionToken(), delivery.Products)
	rg.POST("/products", delivery.parseSessionToken(), delivery.UploadProducts)
}

func (d *adminDelivery) MainPanel(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Admin Panel",
	})
}

func (d *adminDelivery) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func (d *adminDelivery) SignIn(c *gin.Context) {
	c.Request.ParseForm()
	sessionToken, expiresAt, err := d.handler.SignIn(c, c.PostForm("username"), c.PostForm("password"))
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
	c.Redirect(http.StatusFound, "/admin/")
}

func (d *adminDelivery) Products(c *gin.Context) {
	path, _ := d.handler.GetFilePath(c, "products")
	c.HTML(http.StatusOK, "products.html", gin.H{
		"filepath": path.FilePath,
	})
}

func (d *adminDelivery) UploadProducts(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.HTML(http.StatusOK, "products.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	filePath := "assets/uploads/" + file.Filename
	assetsPath := domain.AssetsPath{
		Key:      "products",
		FilePath: filePath,
	}
	if err = d.handler.SaveFilePath(c, assetsPath); err != nil {
		c.HTML(http.StatusOK, "products.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = c.SaveUploadedFile(file, filePath); err != nil {
		assetsPath.FilePath = ""
		d.handler.SaveFilePath(c, assetsPath)
		c.HTML(http.StatusOK, "products.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "products.html", gin.H{
		"message":  "Data berhasil diperbarui",
		"filepath": filePath,
	})
}

func (d *adminDelivery) parseSessionToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken, err := c.Cookie("session_token")
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/admin/login")
		}

		_, err = d.handler.GetSession(c, sessionToken)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/admin/login")
		}
		c.Next()
	}
}
