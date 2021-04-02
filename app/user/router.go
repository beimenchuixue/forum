package user

import "github.com/gin-gonic/gin"

// Router 用户模型相关路由
func Router(g *gin.Engine) {
	//1. 用户路由
	g.POST("/signup", SignupHandler)
	g.POST("/login", LoginHandler)
}
