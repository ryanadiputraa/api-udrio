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
		c.Redirect(http.StatusFound, "/admin/login")
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})
	c.Redirect(http.StatusFound, "/admin/")
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
