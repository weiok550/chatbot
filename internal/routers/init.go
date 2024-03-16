package routers

import (
	"chatbot/internal/config"
	"chatbot/internal/config/global"
	"chatbot/internal/models"
	"chatbot/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)



func checkLoginStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var res models.Response
		cookie, err := c.Request.Cookie(global.UserLoginToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, res.Fail(http.StatusUnauthorized, "not login", nil))
			c.Abort()
			return
		}

		userSrv := new(services.UserSrvMock)

		checkok, userInfo := userSrv.CheckLoginStatus(cookie.Value)
		if !checkok {
			c.JSON(http.StatusUnauthorized, res.Fail(http.StatusUnauthorized, "not login", nil))
			c.Abort()
			return

		}
		c.Set("user", userInfo)

		// 如果用户已登录，继续处理请求
		c.Next()
	}
}

func RegisterRoutes(engine *gin.Engine) {
	config.LoadConfig()

	// 跨域  
	engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", global.Cfg.GetString("app.allowOrigin"))
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})


	// 聊天相关路由
	RegisterChatRoutes(engine)

	// 公共路由(免登陆)
	RegisterPublicRoutes(engine)
}
