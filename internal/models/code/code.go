package code

type ResponseCode int

const (
	//业务全局错误码
	Success ResponseCode = 0
	LoginFail ResponseCode = 100001
	GetLoginUserInfoError ResponseCode = 100002

	
)

var ErrorCodeDesc map[ResponseCode]string = make(map[ResponseCode]string)

func init () {
	ErrorCodeDesc[LoginFail] = "wrong username or password"
	ErrorCodeDesc[GetLoginUserInfoError] = "get login user info error"

}
