package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// SignupHandler 用户注册业务逻辑
func SignupHandler(ctx *gin.Context) {
	// 1. 获取参数
	//signupF := new(SignupForm)
	signupU := new(User)
	err := ctx.ShouldBindJSON(signupU)
	if err != nil {
		fmt.Println(err)
		zap.L().Error("signupHandler 获取参数失败", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"error": gin.H{
				"params": "获取参数失败",
			},
			"data": []interface{}{},
		})
		return
	}
	//2. 校验参数
	ok, err, errMap := SignupValid(signupU)
	if err != nil {
		zap.L().Error("signupHandler 校验参数失败", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"error": gin.H{
				"params": "校验参数失败",
			},
			"data": []interface{}{},
		})
		return
	}
	if !ok {
		zap.L().Warn("signupHandler 校验参数失败", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"code":  1001,
			"error": errMap,
			"data":  []interface{}{},
		})
		return
	}
	fmt.Println(signupU)

	// 注册用户
	err = SignupUser(signupU)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"error": gin.H{
				"signup_error": err.Error(),
			},
			"data": []interface{}{},
		})
		return
	}

	// 生成jwt或者设置cookie
	ctx.JSON(http.StatusOK, gin.H{
		"code":  2000,
		"error": gin.H{},
		"jwt":   "x.x.x",
		"data":  []interface{}{},
	})
}
