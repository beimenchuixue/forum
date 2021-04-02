package middleware

import (
	"forum/utils/jwt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// JwtMiddleware jwt认证中间件
func JwtMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// Authorization: Bearer Token
		authorization := ctx.Request.Header.Get("Authorization")
		// 检查长度
		if len(authorization) == 0 {
			ctx.Abort()
			return
		}
		// 尝试分割
		data := strings.Split(authorization, " ")
		if len(data) != 2 {
			ctx.Abort()
			return
		}
		tokenString := data[1]

		// 检查token
		token := jwt.Toke{}
		payload, err := token.CheckToken(tokenString)
		if err != nil {
			ctx.Abort()
			return
		}

		// 验证是否过期
		if payload.Exp < time.Now().Unix() {
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
