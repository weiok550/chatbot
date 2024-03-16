package interfaces

import "chatbot/internal/models"

type IChatbot interface {
	GetResponse(userId uint32, receiveMsg string) models.ChatResp
}
