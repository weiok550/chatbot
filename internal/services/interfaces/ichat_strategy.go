package interfaces
type IChatStrategy interface {
	ChooseBot(userId uint32, receiveMsg string) IChatbot
}