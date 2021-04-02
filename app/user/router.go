package user

import "github.com/gin-gonic/gin"

func Router(g *gin.RouterGroup) {
	g.POST("/join", SignupHandler)
	g.POST("/login", LoginHandler)
}
