package middleware

import (
	"net/http"

	"github.com/Quanghh2233/blogs/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := utils.GetTokenString(ctx)
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorsResponse("No token found"))
			return
		}
		token, err := utils.ValidateToken(tokenString)
		if token == nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorsResponse(err.Error()))
			return
		}
	}
}
