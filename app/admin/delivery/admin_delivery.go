package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type adminDelivery struct{}

func NewAdminDelivery(rg *gin.RouterGroup) {
	deliver := adminDelivery{}
	rg.GET("/", deliver.MainPanel)
	rg.GET("/login", deliver.Login)
}

func (d *adminDelivery) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func (d *adminDelivery) MainPanel(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Admin Panel",
	})
}
