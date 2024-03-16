package models

import (
	"chatbot/internal/models/code"
	"net/http"
)

type Response struct {
	Code code.ResponseCode `json:"code"`
	Msg  string       `json:"msg"`
	Data any          `json:"data"`
}

func (res *Response) Success(data any) Response {
	res.Code = http.StatusOK
	res.Msg = "ok"
	res.Data = data

	return *res
}

func (res *Response) Fail(code code.ResponseCode, msg string, data any) Response {
	res.Code = code
	res.Msg = msg
	res.Data = data

	return *res
}
