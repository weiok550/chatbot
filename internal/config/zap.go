package config

import (
	"chatbot/internal/config/global"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerConfig struct {
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"maxSize"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxAge     int    `yaml:"maxAge"`
	Compress   bool   `yaml:"compress"`
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

var Logger *zap.Logger

func init() {
	LoadConfig()
	cfg := global.Cfg
	var loggerConfig LoggerConfig

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	global.Cfg.UnmarshalKey("logger", &loggerConfig)
	atomicLevel.SetLevel(getLogLevel(loggerConfig.Level))

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   loggerConfig.Filename,   // 日志文件路径
			MaxSize:    loggerConfig.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: loggerConfig.MaxBackups, // 日志文件最多保存多少个备份
			MaxAge:     loggerConfig.MaxAge,     // 文件最多保存多少天
			Compress:   loggerConfig.Compress,   // 是否压缩
		}),
		atomicLevel, // 日志级别
	)

	// 开启文件及行号
	caller := zap.AddCaller()
	// 开启开发模式，堆栈跟踪
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", cfg.GetString("app.serviceName")))
	// 构造日志
	Logger = zap.New(core, caller, development, filed)
}

func LoadLogger() {
	global.Logger = Logger
}
