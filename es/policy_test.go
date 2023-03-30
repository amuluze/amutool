// Package es
// Date: 2023/3/28 13:29
// Author: Amu
// Description:
package es

import (
	"context"
	"fmt"
	"testing"
)

func TestILMExists(t *testing.T) {
	var esClient = getClient()
	policyName := ".ilm-policy"
	exists, err := esClient.ILMPolicyExists(context.Background(), policyName)
	if err != nil {
		fmt.Printf("error: %#v\n", err)
	}
	fmt.Printf("exists: %#v\n", exists)
}

func TestGetILMPolicy(t *testing.T) {
	var esClient = getClient()
	policyName := ".deprecation-indexing-ilm-policy"
	res, err := esClient.GetILMPolicy(context.Background(), policyName)
	if err != nil {
		fmt.Printf("error: %#v\n", err)
	}
	fmt.Println(res)
}
