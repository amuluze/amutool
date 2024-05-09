// Package tail
// Date: 2024/5/8 10:15
// Author: Amu
// Description:
package tail

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestTail(t *testing.T) {
	tail := NewFileTail("/Users/amu/Desktop/github/amutool/tail/test.log")
	go tail.Tail()
	go func() {
		f, err := os.OpenFile("/Users/amu/Desktop/github/amutool/tail/test.log", os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			t.Error("open file error: ", err)
			return
		}
		defer f.Close()
		for i := 0; i < 10; i++ {
			_, err := f.WriteString("hello world " + fmt.Sprintf("%d", i) + "\r\n")
			if err != nil {
				t.Error("write file error: ", err)
				return
			}
		}
	}()
	go func() {
		for msg := range tail.Msg() {
			fmt.Println("msg: ", msg)
		}
	}()
	time.Sleep(10 * time.Second)
	tail.Close()
	time.Sleep(time.Second)
}
