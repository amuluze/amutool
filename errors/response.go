// Package errors
// Date: 2022/8/26 15:30
// Author: Amu
// Description:
package errors

import "fmt"

// Response 定义响应
type Response struct {
	Code    int    // 错误码
	Message string // 错误消息
	Status  int    // 响应状态码
	ERR     error  // 响应错误
}

func (r *Response) Error() string {
	if r.ERR != nil {
		return r.ERR.Error()
	}
	return r.Message
}

func UnWrapResponse(err error) *Response {
	if v, ok := err.(*Response); ok {
		return v
	}
	return nil
}

func WrapResponse(err error, code, status int, msg string, args ...interface{}) error {
	res := &Response{
		Code:    code,
		Message: fmt.Sprintf(msg, args...),
		ERR:     err,
		Status:  status,
	}
	return res
}

func Wrap400Response(err error, msg string, args ...interface{}) error {
	return WrapResponse(err, 0, 400, msg, args...)
}

func Wrap500Response(err error, msg string, args ...interface{}) error {
	return WrapResponse(err, 0, 500, msg, args...)
}

func NewResponse(code, status int, msg string, args ...interface{}) error {
	res := &Response{
		Code:    code,
		Message: fmt.Sprintf(msg, args...),
		Status:  status,
	}
	return res
}

func New400Response(msg string, args ...interface{}) error {
	return NewResponse(0, 400, msg, args...)
}

func New500Response(msg string, args ...interface{}) error {
	return NewResponse(0, 500, msg, args...)
}
