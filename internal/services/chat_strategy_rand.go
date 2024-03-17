package services

import (
	"chatbot/internal/config/global"
	"chatbot/internal/services/interfaces"
	"math/rand"
	"time"
)

type ChatStrategyRand struct {
}

func (cs ChatStrategyRand) ChooseBot(userId uint32, receiveMsg string) interfaces.IChatbot {
	var chatBot interfaces.IChatbot
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成一个0到100之间的随机整数
	num := randomGenerator.Intn(100)
	gptHitRate := global.Cfg.GetInt("openai.hitRate")

	if num > gptHitRate {
		chatBot = new(ChatbotMock)
	} else {
		chatBot = new(ChatbotGpt)
	}
	return chatBot
}
