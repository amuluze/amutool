// Package tail
// Date: 2024/5/8 10:10
// Author: Amu
// Description:
package tail

import (
	"github.com/nxadm/tail"
)

type Tail interface {
	Tail()
	Msg() <-chan string
	Close()
}

type FileTail struct {
	filepath string
	flag     bool
	msgCh    chan string
	closeCh  chan struct{}
}

func NewFileTail(filepath string) *FileTail {
	return &FileTail{
		filepath: filepath,
		flag:     false,
	}
}

func (t *FileTail) Tail() {
	t.flag = true
	t.msgCh = make(chan string, 10)
	t.closeCh = make(chan struct{})
	tails, err := tail.TailFile(t.filepath, tail.Config{ReOpen: true, Follow: true, Location: &tail.SeekInfo{Offset: 0, Whence: 2}})
	if err != nil {
		t.Close()
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		select {
		case <-t.closeCh:
			return
		case line, ok = <-tails.Lines:
			if !ok {
				continue
			}
			t.msgCh <- line.Text
		}
	}
}

func (t *FileTail) Msg() <-chan string {
	return t.msgCh
}

func (t *FileTail) Close() {
	if t.flag {
		t.flag = false
		close(t.msgCh)
		close(t.closeCh)
	}
}
