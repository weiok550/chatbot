package chat

import (
	"chatbot/internal/database"
	"chatbot/internal/handlers/user"
	"chatbot/internal/models"
	"chatbot/internal/models/code"
	"chatbot/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取历史聊天记录接口
func GetChatRecords(c *gin.Context) {
	res := models.Response{}
	req := models.GetChatRecordsReq{}
	data := models.GetChatRecordsResp{
		List : make([]models.ChatRecord, 0),
		Index: 0,
	}


	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, res.Fail(http.StatusBadRequest, "bad request", nil))
		return
	}

	userInfo, err := user.GetLoginUserInfo(c)

	if err != nil {
		c.JSON(http.StatusOK, res.Fail(code.GetLoginUserInfoError, code.ErrorCodeDesc[code.GetLoginUserInfoError], nil))
		return
	}


	chatRecordService := database.ChatRecordService{}

	// 只查询必要字段
	tmpdata, err := chatRecordService.GetChatRecordsByPage(userInfo.UserID,"id,type,user_id,msg,created_at", 20, req.Start)

	if err != nil {
		c.JSON(http.StatusInternalServerError, res.Fail(http.StatusInternalServerError, "服务器开小差了，请稍后再试～", nil))	
		return
	}

	if len(tmpdata) == 0 {
		c.JSON(http.StatusOK, res.Success(data))
		return
	}

	for _,v := range tmpdata {
		record := models.ChatRecord{
			UserID:     v.UserID,
			ID:         v.ID,
			Type:       v.Type,
			Msg:        v.Msg,
			CreatedAt:  v.CreatedAt,
		}
		data.List = append(data.List, record)
	}

	data.Index = data.List[len(tmpdata)-1].CreatedAt
	c.JSON(http.StatusOK, res.Success(data))
}

// 聊天接口 
func Chat(c *gin.Context) {
	var chatRequest models.ChatReq
	res := models.Response{}

	if err := c.ShouldBind(&chatRequest); err != nil {
		c.JSON(http.StatusBadRequest, res.Fail(http.StatusBadRequest, "bad request", nil))
		return
	}

	userInfo, err := user.GetLoginUserInfo(c)

	if err != nil {
		c.JSON(http.StatusOK, res.Fail(code.GetLoginUserInfoError, code.ErrorCodeDesc[code.GetLoginUserInfoError], nil))
		return
	}

	var cst = new(services.ChatStrategyRand)

	chatBot := cst.ChooseBot(userInfo.UserID, chatRequest.Msg)

	chatResponse := chatBot.GetResponse(userInfo.UserID, chatRequest.Msg)
	c.JSON(http.StatusOK, res.Success(chatResponse))
}

// 其他独立业务触发-主动给用户推送消息
func CheckAndAutoSendMsgToUser() {

	// var activity = new(service.ActivityExample)

	// msg := activity.CheckAndGetShouldAutoSendMessageToUser()

	// Cfg

}
