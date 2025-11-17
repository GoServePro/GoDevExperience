package internal

import (
	"bytes"
	"encoding/gob"
	"reflect"
)

// CalculateSize 近似计算value占用的内存大小（简单场景完全够用）
func CalculateSize(value interface{}) int64 {
	// 先判断是否为nil指针/空接口，直接返回默认值
	if value == nil || isNilPointer(value) {
		return 64 // nil值默认64字节
	}

	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	if err := encoder.Encode(value); err != nil {
		return 64 // 其他编码失败场景也返回默认值
	}
	return int64(buf.Len())
}

// isNilPointer 判断是否为nil指针（避免gob编码panic）
func isNilPointer(value interface{}) bool {
	val := reflect.ValueOf(value)
	// 只处理指针/接口类型
	if val.Kind() != reflect.Ptr && val.Kind() != reflect.Interface {
		return false
	}
	return val.IsNil()
}
