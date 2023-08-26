package server

import (
	"net/http"
	"runtime/debug"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/config"
	"github.com/ryanadiputraa/api-udrio/internal/middleware"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
	"gorm.io/gorm"
)

type Server struct {
	conf    *config.Config
	gin     *gin.Engine
	db      *gorm.DB
	redis   database.Redis
	storage *storage.BucketHandle
}

func NewServer(conf *config.Config, db *gorm.DB, redis database.Redis, storage *storage.BucketHandle) *Server {
	return &Server{
		conf:    conf,
		gin:     gin.Default(),
		db:      db,
		redis:   redis,
		storage: storage,
	}
}

func (s *Server) Run() error {
	if s.conf.Server.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	s.gin.Use(customRecovery())
	s.gin.Use(middleware.CORSMiddleware())

	s.gin.SetTrustedProxies(nil)
	s.gin.LoadHTMLGlob("templates/**/*")
	s.gin.Static("/assets", "./assets")
	s.gin.MaxMultipartMemory = 5 << 20 // max 8 MiB

	s.MapHandlers()

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
