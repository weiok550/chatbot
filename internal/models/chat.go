package models


type ChatReq struct {
	Msg string `form:"msg"`
}

type ChatResp struct {
	Msg string `json:"msg"`
}

type GetChatRecordsReq struct {
	Start uint64 `form:"start"`
}

type ChatRecord struct {
    ID          uint64 `json:"id"`
    Type        uint8  `json:"type"`
    UserID      uint32 `json:"user_id"`
    Msg         string `json:"msg"`
    CreatedAt   uint64 `json:"created_at"`
}

type GetChatRecordsResp struct {
	List []ChatRecord `json:"list"`
	Index uint64  `json:"index"`
}


type ChatType uint8

const (
    ChatTypeBot  uint8 = 0
    ChatTypeUser uint8 = 1

    ChatBotProviderMock uint8 = 1
    ChatBotProviderGpt  uint8 = 2
)

var ChatTypeDesc = map[uint8]string{
    ChatTypeBot:  "机器人",
    ChatTypeUser: "用户",
}

var ChatBotProviderDesc = map[uint8]string{
    ChatBotProviderMock:  "Mock",
    ChatBotProviderGpt: "chatgpt",
}