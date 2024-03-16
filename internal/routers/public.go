package routers

import (
	"chatbot/internal/handlers/user"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(engine *gin.Engine) {
	// 不需要检测登录的路由组
	public := engine.Group("/api/public")
	{
		public.POST("/login", user.Login)
	}
}
