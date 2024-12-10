package route

import (
	"github.com/Quanghh2233/blogs/internal/controller"
	"github.com/Quanghh2233/blogs/internal/middleware"
	"github.com/Quanghh2233/blogs/internal/repository"
	"github.com/Quanghh2233/blogs/internal/service"
	"github.com/Quanghh2233/blogs/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommentRoute(db *gorm.DB, CommentRouter *gin.RouterGroup, logger *logger.Logger) {
	var (
		commentRepository repository.CommentRepo       = repository.NewCommentRepo(db)
		commentService    service.CommentService       = service.NewCommentService(commentRepository)
		commentController controller.CommentController = controller.NewCommentController(commentService, logger)
	)

	CommentRouter.GET("", commentController.All)
	CommentRouter.POST("", middleware.AuthorizeJWT(), commentController.Insert)
}
