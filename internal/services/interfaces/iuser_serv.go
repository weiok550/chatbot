package interfaces

import "chatbot/internal/models"

type IUserServ interface {

	// 登陆检测用户名密码是否正确，正确则返回用户信息
	CheckPassword(username string, password string) (bool, models.UserInfo)

	// 根据用户名获取账户信息
	GetAccountInfoByName(username string) models.AccountInfo

	// 检测用户登陆态
	CheckLoginStatus(token string) (res bool, userInfo *models.UserInfo) 

	// 写入用户登陆态
	WriteLoginStatus(userInfo models.UserInfo) (token string, err error)
}
