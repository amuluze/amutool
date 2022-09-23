// Package router
// Date: 2022/9/23 13:29
// Author: Amu
// Description:
package router

type Router struct {
	Path        string      // api url
	Method      string      // api method
	Description string      // api description
	Deprecated  bool        // 是否弃用
	Request     interface{} // api 请求参数
	Responses   interface{} // api 返回内容
	Header      interface{} // api 请求头
	Cookie      interface{} // api Cookie
	Tags        []string    // api 分组
	HasSecurity bool        // 是否有安全认证
}

func New(path, method, description string, tags []string, options ...Option) *Router {
	r := &Router{
		Path:        path,
		Method:      method,
		Description: description,
		Tags:        tags,
	}

	for _, option := range options {
		option(r)
	}

	return r
}
