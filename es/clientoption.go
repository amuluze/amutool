// Package es
// Date: 2023/12/1 15:51
// Author: Amu
// Description:
package es

type ClientOption func(option *clientOption)

type clientOption struct {
	Addr        string
	Username    string
	Password    string
	Sniff       bool
	Debug       bool
	Healthcheck bool
}

func WithAddr(addr string) ClientOption {
	return func(option *clientOption) {
		option.Addr = addr
	}
}

func WithUsername(username string) ClientOption {
	return func(option *clientOption) {
		option.Username = username
	}
}

func WithPassword(password string) ClientOption {
	return func(option *clientOption) {
		option.Password = password
	}
}

func WithSniff(sniff bool) ClientOption {
	return func(option *clientOption) {
		option.Sniff = sniff
	}
}

func WithDebug(debug bool) ClientOption {
	return func(option *clientOption) {
		option.Debug = debug
	}
}

func WithHealthcheck(check bool) ClientOption {
	return func(option *clientOption) {
		option.Healthcheck = check
	}
}
