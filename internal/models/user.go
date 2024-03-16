package models

type UserLoginReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserConfig struct {
	LoginTokenSecret string        `yaml:"loginTokenSecret"`
	Accounts         []AccountInfo `yaml:"accounts"`
}

type AccountInfo struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	UserID   uint32 `yaml:"user_id"`
}

type UserInfo struct {
	Username string `json:"username"`
	UserID   uint32 `json:"user_id"`
}
