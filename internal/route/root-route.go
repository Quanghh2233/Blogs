package route

import (
	"github.com/Quanghh2233/blogs/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RootRoute(db *gorm.DB, router *gin.Engine, logger *logger.Logger) {
	router.Static("/media", "/media")
	apiRouter := router.Group("/api/v1")
	postRouter := apiRouter.Group("/post")
	PostRoute(db, postRouter, logger)
	commentRouter := apiRouter.Group("/posts/:postId/comments")
	CommentRoute(db, commentRouter, logger)
	categoryRouter := apiRouter.Group("/categories")
	CategoryRoute(db, categoryRouter, logger)
	authRouter := apiRouter.Group("/auth")
	AuthRoute(db, authRouter, logger)
}
