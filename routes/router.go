package routes

import (
	"github.com/dtg-lucifer/go-gorm-gin-json-api/controllers"
	"github.com/dtg-lucifer/go-gorm-gin-json-api/storage"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Store      *storage.Storage
	Engine     *gin.Engine
}

func SetupRoutes(store *storage.Storage, engine *gin.Engine) *Routes {
  // NOTE: health check route
	mainGroup := engine.Group("/api/v1")
  {
    mainGroup.GET("/health", HealthCheck)
  }

  // TODO: manage authentication routes
  authGroup := mainGroup.Group("/auth")
  {
    authGroup.GET("/me", func(ctx *gin.Context) {})
  }

  // NOTE: routes related to posts
  postsGroup := mainGroup.Group("/posts")
  {
    postsGroup.GET("/", func(ctx *gin.Context) { controllers.GetAllPosts(ctx, store.DB) })
    postsGroup.POST("/create", func(ctx *gin.Context) { controllers.CreatePost(ctx, store.DB) })
    postsGroup.GET("/:id", func(ctx *gin.Context) { controllers.GetPostById(ctx, store.DB) })
    postsGroup.DELETE("/:id", func(ctx *gin.Context) { controllers.DeletePost(ctx, store.DB) })
    postsGroup.PUT("/:id", func(ctx *gin.Context) { controllers.UpdatePost(ctx, store.DB) })
  }

	return &Routes{
		Store:      store,
		Engine:     engine,
	}
}
