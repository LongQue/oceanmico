package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oceanmico/common"
)

//鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || tokenString[0:7] != "Bearer " {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "用户验证失败",
			})
			ctx.Abort()
			return
		}
		token, claims, err := common.ParseToken(tokenString[7:])
		//验证token
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "用户验证失败",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user", claims.Username)
		ctx.Next()
	}
}
