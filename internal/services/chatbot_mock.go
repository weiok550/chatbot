package services

import (
	"chatbot/internal/models"
	"chatbot/pkg/utils"
)

type ChatbotMock struct {
}

func (cm ChatbotMock) GetResponse(userId uint32, receiveMsg string) models.ChatResp {
	ChatrecordSrv := ChatrecordSrv{}

	// 用户聊天信息异步入库
	userTime := utils.Now()
	ChatrecordSrv.AsyncWriteChatRecord(userId, receiveMsg, models.ChatTypeUser, models.ChatBotProviderMock, "", userTime)

	res := models.ChatResp{Msg: "It seems like you asked the wrong person. I think you should ask chatgpt"}

	// 机器人聊天信息异步入库
	botTime := utils.Now()
	ChatrecordSrv.AsyncWriteChatRecord(userId, res.Msg, models.ChatTypeBot, models.ChatBotProviderMock, "", botTime)

	return res
}
