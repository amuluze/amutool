// Package tail
// Date: 2024/5/8 10:10
// Author: Amu
// Description:
package tail

import (
	"fmt"
	"github.com/nxadm/tail"
)

type Tail interface {
	Tail()
	Close()
}

type FileTail struct {
	filepath string
	flag     bool
	ch       chan struct{}
}

func NewFileTail(filepath string) *FileTail {
	return &FileTail{
		filepath: filepath,
		flag:     false,
		ch:       make(chan struct{}),
	}
}

func (t *FileTail) Tail() {
	t.flag = true
	tails, err := tail.TailFile(t.filepath, tail.Config{ReOpen: true, Follow: true, Location: &tail.SeekInfo{Offset: 0, Whence: 2}})
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		t.flag = false
		return
	}
	fmt.Println("start tail file")
	var (
		line *tail.Line
		ok   bool
	)
	for {
		select {
		case <-t.ch:
			fmt.Println("tail file close...")
			return
		case line, ok = <-tails.Lines:
			if !ok {
				fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
				continue
			}
			fmt.Println("read line:", line.Text)
		}
	}
}

func (t *FileTail) Close() {
	if t.flag {
		t.flag = false
		t.ch <- struct{}{}
	}
}
