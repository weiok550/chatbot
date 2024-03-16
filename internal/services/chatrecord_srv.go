package services

import (
	"chatbot/internal/config/global"
	"chatbot/internal/database"

	"go.uber.org/zap"
)

type ChatrecordSrv struct {
}

func (cs ChatrecordSrv) AsyncWriteChatRecord(userId uint32, msg string, tp uint8, botProvider uint8, botModel string, createdAt uint64) error {
	// 将机器人聊天信息写入数据库，正确的方式是异步写入，比如先写入mq，然后起服务来多线程消费并写入
	// 异步消费写入场景聊天信息的顺序问题可以使用created_at字段来解决，用户侧也是使用created_at字段排序聊天记录
	// 这里为了方便用例执行 直接同步执行入库操作
	chatRecordService := database.ChatRecordService{}
	data := database.ChatRecord{
		Type:        tp,
		UserID:      userId,
		BotProvider: botProvider,
		BotModel:    botModel,
		Msg:         msg,
		CreatedAt:   createdAt,
	}
	err := chatRecordService.CreateChatRecord(data)
	if err != nil {
		global.Logger.Error("chatrecord写入数据库失败:", zap.Any("data", data))
		// 可以写入消息队列 后续重试 有告警系统则接入相应指标
	}
	return err
}
