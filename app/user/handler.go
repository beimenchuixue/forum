package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ValidUserData 获取并校验用户参数
func ValidUserData(ctx *gin.Context) (u *User, ok bool) {
	// 1. 获取参数
	resp := NewResponse(ctx)
	u = new(User)
	err := ctx.ShouldBindJSON(u)
	if err != nil {
		resp.ErrResponse(http.StatusOK, ErrInvalidParam, err, "ShouldBindJSON 绑定参数失败",
			gin.H{"params": "参数出错"}, nil,
		)
		return nil, false
	}

	//2. 校验参数
	ok, err, errMap := ValidData(u)
	// 处理校验这个过程失败
	if err != nil {
		resp.ErrResponse(http.StatusOK, ErrInvalidParam, err, "ValidData 校验过程失败失败",
			gin.H{"params": "参数出错"}, nil,
		)
		return nil, false
	}

	// 处理参数不符合要求
	if !ok {
		resp.ErrResponse(http.StatusOK, ErrInvalidParam, err, "ValidData 校验参数失败",
			errMap, nil,
		)
		return nil, false
	}
	return u, true
}

// SignupHandler 用户注册业务逻辑
func SignupHandler(ctx *gin.Context) {
	u, ok := ValidUserData(ctx)
	if !ok {
		return
	}

	resp := NewResponse(ctx)
	// 注册用户
	err := SignupUser(u)
	if err != nil {
		resp.ErrResponse(http.StatusOK, ErrUserExist, err, "signupHandler 用户已存在",
			gin.H{"params": GetStatusString(ErrUserExist)}, nil,
		)
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

// LoginHandler 处理用户登录
func LoginHandler(ctx *gin.Context) {
	// 获取并校验参数
	u, ok := ValidUserData(ctx)
	if !ok {
		return
	}

	resp := NewResponse(ctx)
	// 验证登录
	_, ok = Login(u.Username, u.Password)
	if !ok {
		resp.ErrResponse(http.StatusOK, ErrInvalidUsernameOrPassword, nil, "loginHandler 用户登录失败",
			gin.H{"params": GetStatusString(ErrInvalidUsernameOrPassword)}, nil,
		)
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
