package server

import (
	"helloblog/internal/config"
	"helloblog/internal/server/middleware"
	"helloblog/internal/svc"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type HTTPServer struct {
	cfg    config.Config
	engine *gin.Engine
}

func New(cfg config.Config, database *gorm.DB, redisClient *redis.Client) *HTTPServer {
	gin.SetMode(cfg.Server.Mode)

	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery(), middleware.CORS())

	serviceContext := svc.NewServiceContext(cfg, database, redisClient)
	RegisterRoutes(engine, serviceContext)

	return &HTTPServer{
		cfg:    cfg,
		engine: engine,
	}
}

func (s *HTTPServer) Run() error {
	return s.engine.Run(s.cfg.Server.Addr)
}
