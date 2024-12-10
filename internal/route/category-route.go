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

func CategoryRoute(db *gorm.DB, categoryRouter *gin.RouterGroup, logger *logger.Logger) {
	var (
		categoryRepository repository.CategoryRepo       = repository.NewCategoryRepo(db)
		categoryService    service.CategoryService       = service.NewCategoryService(categoryRepository)
		categoryController controller.CategoryController = controller.NewCategoryController(categoryService, logger)
	)
	categoryRouter.GET("", categoryController.All)
	categoryRouter.POST("", middleware.AuthorizeJWT(), categoryController.Insert)
	categoryRouter.PUT("/:categoryId", middleware.AuthorizeJWT(), categoryController.Update)
	categoryRouter.DELETE("/:categoryId", middleware.AuthorizeJWT(), categoryController.DeleteById)
}
