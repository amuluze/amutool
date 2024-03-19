// Package es
// Date: 2023/3/30 15:12
// Author: Amu
// Description:
package es

type Model interface {
	GetIndexName() string
	GetId() string
}
