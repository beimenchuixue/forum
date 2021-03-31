package logger

import (
	"forum/settings"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
)

func init() {
	// 日志编码器、输出地址、日志水平配置
	zaoCore := zapcore.NewCore(setEncoder(), setOutput(), setLogLevel())
	// 创建一个日志记录器对象
	l := zap.New(zaoCore, zap.AddCaller())
	// 替换zap中全局日志记录器
	zap.ReplaceGlobals(l)
}

// setEncoder 设置日志编码器格式
func setEncoder() zapcore.Encoder {
	// 日志记录器配置
	pCfg := zap.NewProductionEncoderConfig()
	// a. 日期格式
	pCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	// b. info -> INFO
	pCfg.EncodeLevel = zapcore.CapitalLevelEncoder

	// 设置json作为编码器格式
	jsonEncoder := zapcore.NewJSONEncoder(pCfg)
	return jsonEncoder
}

// setOutput 设置日志输出位置
func setOutput() (sync zapcore.WriteSyncer) {
	// 如果是debug则将输出到终端
	// 如果是release则将输出到文件
	switch settings.Conf.Mode {
	case "debug":
		sync = zapcore.AddSync(os.Stdout)
	case "release":
		// 使用lumberjack作为自动日志切割工具
		l := &lumberjack.Logger{
			Filename:   path.Join(settings.Conf.LogConfig.Dir, settings.Conf.LogConfig.Filename),
			MaxSize:    settings.Conf.LogConfig.MaxSize,
			MaxAge:     settings.Conf.LogConfig.MaxAge,
			MaxBackups: settings.Conf.LogConfig.MaxBackup,
			LocalTime:  true,
		}
		sync = zapcore.AddSync(l)
	}

	return sync
}

// setLogLevel 设置日志级别
func setLogLevel() zapcore.LevelEnabler {
	switch level := settings.Conf.LogConfig.Level; level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	default:
		return zap.ErrorLevel
	}
}
