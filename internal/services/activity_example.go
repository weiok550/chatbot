package services

type ActivityExample struct {
}

func (ae *ActivityExample) CheckAndGetShouldAutoSendMessageToUser(userId int32) string {
	return "this is an example activity notice msg"
}
