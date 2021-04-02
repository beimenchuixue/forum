package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 定义返回的信息结构

/*
{
	"code": int,
	"error": map[string]string,
	"data": map[string]interface{}
}
*/

// JsonData 响应请求的json格式
type JsonData struct {
	Code  StatusCode             `json:"code"`
	Error interface{}            `json:"error"`
	Data  map[string]interface{} `json:"data"`
}

func NewJsonData(code StatusCode, err interface{}, data map[string]interface{}) *JsonData {
	return &JsonData{
		Code:  code,
		Error: err,
		Data:  data,
	}
}

// Response 响应总共分三类，由用户参数引起的错误，系统内部逻辑错误，正确响应
type Response struct {
	ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		ctx: ctx,
	}
}

// ErrParamResponse 用户传递的参数错误响应
// httpCode http状态码
// statusCode 响应状态码
// err 内部处理错误，需要打印到日志
// msg 错误信息
// logicErr 业务逻辑错误
// data 是返回的数据
func (r *Response) ErrResponse(httpCode int, statusCode StatusCode, err error, msg string, logicErr gin.H, data map[string]interface{}) {
	zap.L().Error(msg, zap.Error(err))
	jsonD := NewJsonData(statusCode, logicErr, data)
	r.ctx.JSON(httpCode, jsonD)
}

// CorrectResponse 正确响应
func (r *Response) CorrectResponse(httpCode int, statusCode StatusCode, data map[string]interface{}) {
	jsonD := NewJsonData(statusCode, nil, data)
	r.ctx.JSON(httpCode, jsonD)
}
