// Package errors
// Date: 2023/11/23 14:55
// Author: Amu
// Description:
package errors

import "github.com/pkg/errors"

const (
	InternalServeError    = "internal server error"
	InvalidToken          = "invalid token"
	InvalidUserOrPassword = "invalid username or password"
	MethodNotAllow        = "method not allowed"
	NotFound              = "not found"
	TooManyRequests       = "too many requests"
	Forbidden             = "forbidden"
	BadRequest            = "bad request"
)

// Define alias
var (
	Is           = errors.Is
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
)

var (
	ErrInvalidToken    = NewError(401, InvalidToken)
	ErrForbidden       = NewError(403, Forbidden)
	ErrNotFound        = NewError(404, NotFound)
	ErrMethodNotAllow  = NewError(405, MethodNotAllow)
	ErrTooManyRequests = NewError(429, TooManyRequests)
	ErrInternalServer  = New500Error(InternalServeError)
	ErrBadRequest      = New400Error(BadRequest)
	ErrInvalidAccount  = New400Error(InvalidUserOrPassword)
)

// Error 定义响应错误
type Error struct {
	Message string // 错误消息
	Status  int    // 响应状态码
	ERR     error  // 响应错误
}

func (r *Error) Error() string {
	if r.ERR != nil {
		return r.ERR.Error()
	}
	return r.Message
}

func (r *Error) CodeStatus() int {
	return r.Status
}

func NewError(status int, msg string) *Error {
	return &Error{Status: status, Message: msg, ERR: errors.New(msg)}
}

func New400Error(msg string) error {
	return &Error{Status: 400, Message: msg, ERR: errors.New(msg)}
}

func New500Error(msg string) error {
	return &Error{Status: 500, Message: msg, ERR: errors.New(msg)}
}

func UnWrapError(err error) *Error {
	var v *Error
	if errors.As(err, &v) {
		return v
	}
	return nil
}

func WrapError(err error, status int, msg string) error {
	res := &Error{
		Message: msg,
		ERR:     err,
		Status:  status,
	}
	return res
}
