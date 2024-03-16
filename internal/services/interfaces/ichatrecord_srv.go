package interfaces

type IChatrecordSrv interface {
	AsyncWriteChatRecord(userId uint32, msg string, tp uint8, botProvider uint8, botModel string, createdAt uint64) error
}