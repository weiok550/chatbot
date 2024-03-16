package utils

import (
	"crypto/md5"
	"crypto/subtle"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 生成n位随机数(由字母数字组成)
func GenerateRandomString(length int) string {
	rand.New(rand.NewSource(time.Now().UnixMicro()))

	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < 8; i++ {
		result += string(charSet[rand.Intn(len(charSet))])
	}

	return result
}

func GenerateLoginToken(userID uint32, secret string) string {
	// 将用户id和当前时间戳ms转换为字符串
	userIDStr := strconv.Itoa(int(userID))
	timestamp := Now()

	// 计算 MD5 散列值
	hash := md5.Sum([]byte(userIDStr + "_" + strconv.FormatUint(timestamp, 10) + "_" + secret))

	// 将散列值转换为十六进制字符串
	hashString := hex.EncodeToString(hash[:])

	// 从第8位截取hashString到末尾然后拼接上key
	hashString = hashString[8:]

	// 拼接用户id、时间戳和密钥，并返回
	return userIDStr + "_" + strconv.FormatUint(timestamp, 10) + "_" + hashString
}

func CheckLoginToken(token string, secret string) (bool, int32, int64) {
	// 根据generateLoginToken逻辑反向校验token的合法性
	tokenParts := strings.Split(token, "_")
	if len(tokenParts) != 3 {
		return false, 0, 0
	}

	userId, err := strconv.Atoi(tokenParts[0])
	if err != nil {
		return false, 0, 0
	}

	timestamp, err := strconv.ParseInt(tokenParts[1], 10, 64)
	if err != nil {
		return false, 0, 0
	}

	hash := md5.Sum([]byte(tokenParts[0] + "_" + tokenParts[1] + "_" + secret))
	hashString := hex.EncodeToString(hash[:])
	expectedToken := tokenParts[0] + "_" + tokenParts[1] + "_" + hashString[8:]

	if subtle.ConstantTimeCompare([]byte(token), []byte(expectedToken)) == 1 {
		return true, int32(userId), timestamp
	}

	return false, 0, 0

}
