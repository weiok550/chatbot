package routers

import (
	"chatbot/internal/handlers/chat"

	"github.com/gin-gonic/gin"
)

func RegisterChatRoutes(engine *gin.Engine) {
	// 需要检测登录的路由组
	authorized := engine.Group("/api/chat")
	authorized.Use(checkLoginStatus())
	{
		authorized.GET("/getChatRecords", chat.GetChatRecords)
		authorized.POST("/chat", chat.Chat)
	}

}
