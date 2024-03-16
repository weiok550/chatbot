package services

import (
	"chatbot/internal/config/global"
	"chatbot/internal/models"
	"chatbot/internal/redis"
	"chatbot/pkg/utils"
	"encoding/json"
	"time"
)

type UserSrvMock struct {
}


func (usm *UserSrvMock) CheckPassword(username string, password string) (bool, models.UserInfo) {
	var res models.UserInfo
	accountInfo := usm.GetAccountInfoByName(username)
	if accountInfo.Password == password {
		res.UserID = accountInfo.UserID
		res.Username = accountInfo.Username
		return true, res
	}else {
		return false, res
	}
}

func (usm *UserSrvMock) GetAccountInfoByName(username string) models.AccountInfo {
	var res models.AccountInfo
	var userConfig models.UserConfig
	global.Cfg.UnmarshalKey("user", &userConfig)
	for _,v := range userConfig.Accounts {
		if v.Username == username {
			return v
		}
	}
	return res
}

func (usm *UserSrvMock) CheckLoginStatus(token string) (res bool, userInfo *models.UserInfo) {
	var userConfig models.UserConfig
	global.Cfg.UnmarshalKey("user", &userConfig)

	if checkok,_,_ := utils.CheckLoginToken(token, userConfig.LoginTokenSecret); !checkok{
		return false, nil
	}

	// 优先从本地缓存中读取用户信息 todo

	// 从redis读取用户信息
	userInfoStr,err := redis.GetKeyValue("user", token)
	if err != nil {
		return false, nil
	}

	if res := json.Unmarshal([]byte(userInfoStr), &userInfo); res != nil {
		return false, nil
	}

	// 登陆用户信息写入本地缓存 todo
	return true, userInfo
}

func (usm *UserSrvMock) WriteLoginStatus(userInfo models.UserInfo) (token string, err error) {
	var userConfig models.UserConfig
	global.Cfg.UnmarshalKey("user", &userConfig)

	token = utils.GenerateLoginToken(userInfo.UserID, userConfig.LoginTokenSecret)

	// 将token作为key，用户信息作为value 写入redis
	jsonData, _ := json.Marshal(userInfo)
	err = redis.SetKeyWithExpiration("user", token, string(jsonData) , 24 * time.Hour)

	// 登陆用户信息写入本地缓存 todo
	
	if err != nil {
		return "", err
	}
	return token, err
}

