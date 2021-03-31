package web

import (
	_ "bbs/conn" // 执行初始化工作
	_ "bbs/logger"
	"bbs/middleware"
	"bbs/settings"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// App 整个webApp对象
type App struct {
	Name    string
	Version string
}

// Run 整个app对象入口
func (a *App) Run() {

	// 数据迁移
	if err := autoMigrate(); err != nil {
		zap.L().Error("数据迁移失败", zap.Error(err))
		panic(err)
	}

	// 设置验证器提示消息
	a.setValidatorErrorMsg()

	// 设置gin项目模式
	a.setGinMode()
	// web项目
	engine := gin.New()
	// 启动gin日志中间件和gin出错恢复中间件
	engine.Use(middleware.GinLogger(zap.L()), middleware.GinRecovery(zap.L(), true))
	// 注册路由
	RouterManager(engine)
	// 运行项目
	err := engine.Run(a.getListenAddr())
	if err != nil {
		panic(err)
	}
}

// getListenAddr 得到一个监听地址
func (a *App) getListenAddr() string {
	bind := settings.Conf.Bind
	port := settings.Conf.Port
	if bind == "*" {
		return fmt.Sprintf("%s:%d", "", port)
	}
	return fmt.Sprintf("%s:%d", bind, port)
}

// setValidatorErrorMsg全局设置验证器信息
func (a *App) setValidatorErrorMsg() {
	// 修改对应约束英文错误为中文
	validation.SetDefaultMessage(map[string]string{
		"Required": "不能为空",
	})
}

// setGinMode 设置gin的模式
func (a *App) setGinMode() {
	switch settings.Conf.Mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
}

// NewApp 实例化整个web对象
func NewApp() *App {
	return &App{
		Name:    settings.Conf.Name,
		Version: settings.Conf.Version,
	}
}
