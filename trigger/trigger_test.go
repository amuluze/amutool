// Package trigger
// Date: 2023/7/19 17:04
// Author: Amu
// Description:
package trigger

import (
	"fmt"
	"testing"
)

func testMessage(message interface{}) error {
	fmt.Printf("hello %s\n", message)
	return nil
}

type UserInfo struct {
	Name string
	Age  int
}

func updateUser(userInfo interface{}) error {
	info := userInfo.(UserInfo)
	fmt.Println(info)
	return nil
}

func TestTrigger(t *testing.T) {
	trigger := NewTrigger()
	err := trigger.RegisterEvent("testMessage", testMessage)
	if err != nil {
		fmt.Printf("register event %s error: %v\n", "testMessage", err)
	}
	err = trigger.RegisterEvent("updateUser", updateUser)
	if err != nil {
		fmt.Printf("register event %s error: %v\n", "updateUser", err)
	}

	err = trigger.CallEvent("testMessage", "golang")
	if err != nil {
		fmt.Printf("call event error: %v\n", err)
	}
	userInfo := UserInfo{
		Name: "jack",
		Age:  13,
	}
	err = trigger.CallEvent("updateUser", userInfo)
	if err != nil {
		fmt.Printf("call updateUser error: %v\n", err)
	}
	fmt.Println(trigger.Events())
}
