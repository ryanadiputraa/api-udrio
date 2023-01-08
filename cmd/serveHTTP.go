package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
	"github.com/spf13/viper"
)

func serveHTTP() {
	_ = database.GetConnection()

	r := gin.Default()
	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	// Setup server port & handler
	port := viper.GetString("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	r.Run(fmt.Sprintf(":%s", port))
}
