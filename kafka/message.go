// Package kafka
// Date: 2023/11/30 15:05
// Author: Amu
// Description:
package kafka

type Message interface {
	Encode() ([]byte, error)
	Length() int
}
