// Package helper 帮助函数，时间、数组的通用处理
package helper

import (
	"fmt"
	"go-stress-testing/dynamic"
	"reflect"
	"time"
)

type FuncCollection map[string]reflect.Value

// DiffNano 时间差，纳秒
func DiffNano(startTime time.Time) (diff int64) {
	diff = int64(time.Since(startTime))
	return
}

// InArrayStr 判断字符串是否在数组内
func InArrayStr(str string, arr []string) (inArray bool) {
	for _, s := range arr {
		if s == str {
			inArray = true
			break
		}
	}
	return
}

// CallFunc 动态调用方法
func CallFunc(funcName string, args ... interface{}) (result []reflect.Value, err error) {
	var router dynamic.Routers
	FuncMap := make(FuncCollection, 0)
	reflectValueOf := reflect.ValueOf(&router)
	rft := reflectValueOf.Type()
	funcNum := reflectValueOf.NumMethod()
	for i := 0; i < funcNum; i ++ {
		methodName := rft.Method(i).Name
		FuncMap[methodName] = reflectValueOf.Method(i)
	}
	parameter := make([]reflect.Value, len(args))
	for k, arg := range args {
		parameter[k] = reflect.ValueOf(arg)
	}
	result = FuncMap[funcName].Call(parameter)
	return
}

// Capitalize 首字母转大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}