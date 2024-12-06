package controller

import (
	"github.com/Quanghh2233/blogs/internal/service"
	"github.com/Quanghh2233/blogs/pkg/logger"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
	VerifyToken(context *gin.Context)
	RefreshToken(context *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.jwtService
	logger      *logger.Logger
}

func NewAuthController(authService service.AuthService, jwtService service.JwtService, logger *logger.Logger) *authController {
	return &authController{authService: authService, jwtService: jwtService, logger: logger}
}
