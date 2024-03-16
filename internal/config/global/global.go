package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

/**
 * 项目级全局变量
 */
var (
	Logger *zap.Logger  // 日志
	Cfg    *viper.Viper // 配置
	UserLoginToken string = "cult"	// chatbot user login token
)
