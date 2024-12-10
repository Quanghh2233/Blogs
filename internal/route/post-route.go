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

func PostRoute(db *gorm.DB, postRouter *gin.RouterGroup, logger *logger.Logger) {
	var (
		postRepository repository.PostRepo       = repository.NewPostRepo(db)
		postService    service.PostService       = service.NewPostService(postRepository)
		postController controller.PostController = controller.NewPostController(postService, logger)
	)
	postRouter.GET("", postController.All)
	postRouter.GET("/:postId", postController.FindByID)
	postRouter.POST("", middleware.AuthorizeJWT(), postController.Insert)
	postRouter.PUT("/:postId", middleware.AuthorizeJWT(), postController.Update)
	postRouter.DELETE("/:postId", middleware.AuthorizeJWT(), postController.DeleteById)

}
