// Package convert
// Date: 2022/9/7 01:04
// Author: Amu
// Description:
package convert

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// AsString any to string
func AsString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case time.Time:
		return v.Format("2006-01-02 15:04:05")
	case bool:
		return strconv.FormatBool(v)
	default:
		{
			b, _ := json.Marshal(v)
			return string(b)
		}
	}
}

// ToInt64 any to int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	case string:
		d, err = strconv.ParseInt(val.String(), 10, 64)
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}

// BytesToInt64 byte to int64
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// Int64ToBytes int64 转 []byte
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// StringToInt string to int
func StringToInt(value string) (i int) {
	i, _ = strconv.Atoi(value)
	return i
}

// IntToInt64 int to int64
func IntToInt64(value int) int64 {
	i, _ := strconv.ParseInt(string(rune(value)), 10, 64)
	return i
}
