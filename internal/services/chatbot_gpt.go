package services

import (
	"bytes"
	"chatbot/internal/config/global"
	"chatbot/internal/models"
	"chatbot/pkg/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ChatbotGpt struct {
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Choice struct {
	Index        int      `json:"index"`
	Message      Message  `json:"message"`
	Logprobs     []string `json:"logprobs,omitempty"`
	FinishReason string   `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type OpenAIResponse struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	SystemFingerprint string   `json:"system_fingerprint"`
}

var ChatModel = "gpt-3.5-turbo"

func (cg ChatbotGpt) GetResponse(userId uint32, receiveMsg string) models.ChatResp {

	ChatrecordSrv := ChatrecordSrv{}

	// 用户聊条信息异步入库
	userTime := utils.Now()
	ChatrecordSrv.AsyncWriteChatRecord(userId, receiveMsg, models.ChatTypeUser, models.ChatBotProviderGpt, "", userTime)

	// Parse request body
	var request = OpenAIRequest{
		Model: ChatModel,
		Messages: []Message{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: receiveMsg},
		},
	}

	// Call OpenAI API
	response, err := callOpenAIAPI(request)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("call openai api fail, error:%s", err.Error()))
		return models.ChatResp{Msg: ""}
	}

	// 机器人聊天信息异步入库
	botTime := utils.Now()
	ChatrecordSrv.AsyncWriteChatRecord(userId, response, models.ChatTypeBot, models.ChatBotProviderGpt, ChatModel, botTime)

	return models.ChatResp{Msg: response}

}

func callOpenAIAPI(request OpenAIRequest) (string, error) {
	openAIKey := global.Cfg.GetString("openai.secretKey")

	// Prepare request body
	requestBody, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	// Create HTTP request
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", openAIKey))

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse JSON response
	var response OpenAIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	// Extract and return text from response
	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response text found")
}
