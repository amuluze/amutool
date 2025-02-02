// Package jsonrpc
// Date: 2023/4/21 11:38
// Author: Amu
// Description:
package jsonrpc

import (
	"github.com/gorilla/rpc/v2"
)

type Server struct {
	*rpc.Server
}

func NewServer() (*Server, error) {
	return &Server{rpc.NewServer()}, nil
}
