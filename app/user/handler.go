package user

import (
	"forum/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginAndSignupCorrectResponse(resp *Response, u *User) {
	token := jwt.Toke{}
	tokenStr, err := token.GetToken(u.UserId)
	if err != nil {
		resp.ErrResponse(http.StatusOK, ErrServiceBusy, err, "GetToken json解析错误",
			gin.H{"params": GetStatusString(ErrServiceBusy)}, nil,
		)
		return
	}
	// 返回正确的响应
	resp.CorrectResponse(http.StatusOK, LonginSuccess, gin.H{"token": tokenStr})
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

	// // 生成jwt, 返回token数据
	LoginAndSignupCorrectResponse(resp, u)
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
	// 生成jwt, 返回token数据
	LoginAndSignupCorrectResponse(resp, u)
}
