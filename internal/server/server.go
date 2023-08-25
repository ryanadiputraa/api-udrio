package server

import (
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/config"
	"github.com/ryanadiputraa/api-udrio/internal/middleware"
)

type Server struct {
	conf *config.Config
	gin  *gin.Engine
}

func NewServer(conf *config.Config) *Server {
	return &Server{
		conf: conf,
		gin:  gin.Default(),
	}
}

func (s *Server) Run() error {
	if s.conf.Server.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	s.gin.SetTrustedProxies(nil)
	s.gin.Use(customRecovery())
	s.gin.Use(middleware.CORSMiddleware())

	port := strconv.Itoa(s.conf.Server.Port)
	return s.gin.Run(":" + port)
}

func customRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "something went wrong, please try again later",
				})
			}
		}()

		ctx.Next()
	}
}
