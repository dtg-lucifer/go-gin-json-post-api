package routes

import (
	"github.com/dtg-lucifer/go-gorm-gin-json-api/storage"
	"github.com/gin-gonic/gin"
)

type Routes struct {
  Store *storage.Storage
  Engine *gin.Engine
  RouteGroup *gin.RouterGroup
}

func SetupRoutes(store *storage.Storage, engine *gin.Engine) *Routes {
  rgroup := engine.Group("/api/v1")

  rgroup.GET("/health", HealthCheck)

  routes := &Routes{
    Store: store,
    Engine: engine,
    RouteGroup: rgroup,
  }
  return routes
}
