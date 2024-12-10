package route

import (
	"github.com/Quanghh2233/blogs/internal/controller"
	"github.com/Quanghh2233/blogs/internal/repository"
	"github.com/Quanghh2233/blogs/internal/service"
	"github.com/Quanghh2233/blogs/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoute(db *gorm.DB, authRouter *gin.RouterGroup, logger *logger.Logger) {
	var (
		jwtService     service.JwtService        = service.NewJwtService()
		authRepository repository.AuhtRepo       = repository.NewAuthRepo(db)
		authService    service.AuthService       = service.NewAuthService(authRepository)
		authController controller.AuthController = controller.NewAuthController(authService, jwtService, logger)
	)
	authRouter.POST("/login", authController.Login)
	authRouter.POST("/signup", authController.Register)
	authRouter.POST("/token/verify", authController.VerifyToken)
	authRouter.POST("/token/refresh", authController.RefreshToken)
}
