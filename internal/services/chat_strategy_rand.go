package services

import (
	"chatbot/internal/services/interfaces"
	"math/rand"
	"time"
)

type ChatStrategyRand struct {
}

func (cs ChatStrategyRand) ChooseBot(userId uint32, receiveMsg string) interfaces.IChatbot {
	var chatBot interfaces.IChatbot
	// 创建一个新的随机数生成器
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成一个0到9之间的随机整数
	num := randomGenerator.Intn(10)

	if num < 4 {
		chatBot = new(ChatbotMock)
	} else {
		chatBot = new(ChatbotGpt)
	}
	return chatBot
}
