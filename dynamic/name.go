package dynamic

import (
	"math/rand"
	"time"
)

// Name 构造Name参数
func (router *Routers) Name(args ...interface{}) string {
	return RandString(5)
}

// RandString 随机字符串
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
