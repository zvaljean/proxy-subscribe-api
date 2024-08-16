package utils

import (
	"math/rand"
	"time"
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
