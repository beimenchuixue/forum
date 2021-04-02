package web

import (
	"forum/app/user"
	"forum/middleware"
	"github.com/gin-gonic/gin"
)

// Manager 路由管理者，路由入口，完成路由注册
func RouterManager(router *gin.Engine) {
	router.GET("/ping", middleware.JwtMiddleware(), Index)

	user.Router(router)
}
