package util

import (
	"reflect"
)

func CheckTypeByReflectZero(arg interface{}) bool {
	if reflect.ValueOf(arg).IsZero() { //利用反射直接判空，基础数据类型用isZero
		return reflect.ValueOf(arg).IsValid()
	}
	return false
}
