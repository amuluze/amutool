// Package errors
// Date: 2022/8/26 15:30
// Author: Amu
// Description:
package errors

import "fmt"

// Response 定义响应错误
type Response struct {
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

func WrapResponse(err error, status int, msg string, args ...interface{}) error {
	res := &Response{
		Message: fmt.Sprintf(msg, args...),
		ERR:     err,
		Status:  status,
	}
	return res
}

func Wrap400Response(err error) error {
	return WrapResponse(err, 400, InvalidParameter)
}

func Wrap500Response(err error) error {
	return WrapResponse(err, 500, InternalServeError)
}

func NewResponse(status int, msg string, args ...interface{}) error {
	res := &Response{
		Message: fmt.Sprintf(msg, args...),
		Status:  status,
	}
	return res
}

func New400Response(msg string, args ...interface{}) error {
	return NewResponse(400, msg, args...)
}

func New500Response(msg string, args ...interface{}) error {
	return NewResponse(500, msg, args...)
}
