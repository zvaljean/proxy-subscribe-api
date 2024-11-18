package utils

import (
	"encoding/base64"
	"math/rand"
	"time"
	"zvaljean/proxy-subscribe-api/pkg/log"
)

// 生成指定长度的随机字符串
func GenerateRandomString(length int) string {
	rand.NewSource(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func GenerateRandomInt(scope int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(scope)
}

func Base64toStr(data string) string {

	str, err := base64.StdEncoding.DecodeString(data)
	if log.ErrorCheck(err, "decode error") {
		return ""
	}
	return string(str)
}

func StrtoBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
