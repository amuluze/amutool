// Package jsonrpc
// Date: 2023/4/21 11:32
// Author: Amu
// Description:
package jsonrpc

import (
	"github.com/gorilla/rpc/v2"
)

type JSONRPCService struct {
	codecs map[string]rpc.Codec
}
