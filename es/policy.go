// Package es
// Date: 2023/3/27 19:14
// Author: Amu
// Description: 索引声明周期管理策略
package es

import (
	"context"
	"fmt"
)

const (
	CreatePolicyRetry = 50
	PolicyFilePrefix  = "ilm_"
	PolicyFileSuffix  = ".json"
)

func (c *Client) ILMPolicyExists(ctx context.Context, policyName string) (bool, error) {
	res, err := c.XPackIlmGetLifecycle().Pretty(true).Human(true).Policy(policyName).Do(ctx)
	if err != nil {
		return false, err
	}
	if _, ok := res[policyName]; ok {
		return true, nil
	}
	return false, nil
}

func (c *Client) PutILMPolicy(ctx context.Context, policyName string, policyPath string) error {
	fileName := fmt.Sprint(policyPath, "/", PolicyFilePrefix, policyName, PolicyFileSuffix)
	fmt.Printf("policy file name: %v\n", fileName)
	bodyString, err := c.ReadFile(fileName)
	if err != nil {
		fmt.Printf("read file error: %#v", err)
		return err
	}
	fmt.Printf("body string: %#v\n", bodyString)
	res, err := c.XPackIlmPutLifecycle().Policy(policyName).BodyString(bodyString).Do(context.TODO())
	fmt.Printf(">>>>res: %v, err: %v\n", res, err)
	if err != nil || res.Acknowledged {
		return err
	}
	fmt.Printf("res: %#v\n", res)
	return nil
}

func (c *Client) GetILMPolicy(ctx context.Context, policyName string) (map[string]interface{}, error) {
	res, err := c.XPackIlmGetLifecycle().Pretty(true).Human(true).Policy(policyName).Do(ctx)
	if err != nil {
		return nil, err
	}
	return res[policyName].Policy, nil
}

func (c *Client) DeleteILMPolicy(ctx context.Context, policyName string) error {
	res, err := c.XPackIlmDeleteLifecycle().Policy(policyName).Human(true).Do(ctx)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
