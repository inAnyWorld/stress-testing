package dynamic

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

// Age 构造age参数
func (router *Routers) Age(args ...interface{}) string {
	return strconv.Itoa(int(RandInt64(1,101)))
}

// RandInt64 返回随机数字
func RandInt64(min,max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ :=rand.Int(rand.Reader,maxBigInt)
	if i.Int64() < min {
		RandInt64(min,max)
	}
	return i.Int64()
}