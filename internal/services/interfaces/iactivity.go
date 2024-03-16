package interfaces

type IActivity interface {

	CheckAndGetShouldAutoSendMessageToUser(userId int32) string

}