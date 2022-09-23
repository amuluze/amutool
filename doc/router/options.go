// Package router
// Date: 2022/9/23 13:29
// Author: Amu
// Description:
package router

type Option func(router *Router)

// Requests 设置请求参数
func Requests(request interface{}) Option {
	return func(router *Router) {
		router.Request = request
	}
}

// Deprecated mark api is deprecated
func Deprecated() Option {
	return func(router *Router) {
		router.Deprecated = true
	}
}

// Responses 设置返回内容
func Responses(response interface{}) Option {
	return func(router *Router) {
		router.Responses = response
	}
}

// Headers 设置请求头
func Headers(header interface{}) Option {
	return func(router *Router) {
		router.Header = header
	}
}

// Cookies 设置 cookie
func Cookies(cookie interface{}) Option {
	return func(router *Router) {
		router.Cookie = cookie
	}
}

func HasSecurity(hasSecurity bool) Option {
	return func(router *Router) {
		router.HasSecurity = hasSecurity
	}
}
