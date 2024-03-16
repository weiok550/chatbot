package user

import (
	"chatbot/internal/config"
	"chatbot/internal/config/global"
	"chatbot/internal/models"
	"chatbot/internal/models/code"
	"chatbot/internal/services"
	"chatbot/internal/services/interfaces"
	"chatbot/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userSrv interfaces.IUserServ

func checkLoginStatus(userId int32) {

}

func GetLoginUserInfo(c *gin.Context) (userInfo *models.UserInfo, err error) {
	// 从上下文中获取用户信息
	userInfoInterface, exists := c.Get("user")
	if !exists {
		global.Logger.Fatal("未获取到登录用户信息" + err.Error())
		return nil, err
	}
	userInfo, ok := userInfoInterface.(*models.UserInfo)
	if !ok {
		global.Logger.Fatal("未获取到登录用户信息" + err.Error())
		return nil, err
	}


	return userInfo, err
}

func Login(c *gin.Context) {
	var req models.UserLoginReq
	res := models.Response{}
	if err := c.ShouldBind(&req); err != nil || utils.StringIsEmpty(req.Username) || utils.StringIsEmpty(req.Password) {
		c.JSON(http.StatusBadRequest, res.Fail(http.StatusBadRequest, "bad request", nil))
		return
	}

	userSrv = new(services.UserSrvMock)
	ok, userInfo := userSrv.CheckPassword(req.Username, req.Password)

	if !ok {
		c.JSON(http.StatusOK, res.Fail(code.LoginFail, code.ErrorCodeDesc[code.LoginFail], nil))
		return
	}

	token, err := userSrv.WriteLoginStatus(userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.Fail(http.StatusInternalServerError, "服务器开小差了，请稍后再试～", nil))
		return
	}
	var appSettings config.AppSettings
	global.Cfg.UnmarshalKey("app", &appSettings)

	c.SetCookie(global.UserLoginToken, token, int(appSettings.LoginCookieDuration), "/", appSettings.CookieDomain, false, true)

	c.JSON(http.StatusOK, res.Success(userInfo))
}
