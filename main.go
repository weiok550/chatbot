package main

import (
	"chatbot/internal/config"
	_ "chatbot/internal/config/global"
	"chatbot/internal/routers"
	"chatbot/internal/scripts"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 加载配置
	config.LoadConfig()

	// 初始化日志组件
	config.LoadLogger()

	// 注册路由
	routers.RegisterRoutes(router)

	// 定时任务注册
	scripts.RegisterCronJob()

	// endless无缝重启
	endless.ListenAndServe(":8080", router)

}
