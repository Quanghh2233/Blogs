package main

import (
	"os"

	"github.com/Quanghh2233/blogs/adapter"
	"github.com/Quanghh2233/blogs/docs"
	"github.com/Quanghh2233/blogs/internal/route"
	"github.com/Quanghh2233/blogs/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var db *gorm.DB = adapter.ConnectWithDB()

// @title           Gin Blog Api
// @version         1.0
// @description     A Blog API in Go using Gin framework.

// @contact.name   Hoàng Huy Quang
// @contact.url
// @contact.email  quanghoanghuy33@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1

func main() {
	defer adapter.CloseDbConnection(db)
	logger := logger.NewLogger()
	router := gin.Default()
	route.RootRoute(db, router, logger)
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + os.Getenv("APP_PORT"))
}
