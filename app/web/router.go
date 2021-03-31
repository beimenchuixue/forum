package web

import (
	"forum/app/user"
	"github.com/gin-gonic/gin"
)

// Manager 路由管理者，路由入口，完成路由注册
func RouterManager(router *gin.Engine) {
	router.GET("/", Index)

	//1. 用户路由
	userG := router.Group("/users")
	user.Router(userG)
}
