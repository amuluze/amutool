// Package es
// Date: 2023/3/27 19:14
// Author: Amu
// Description: 索引声明周期管理策略
package es

import (
	"context"
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

func (c *Client) PutILMPolicy(ctx context.Context, policyName string, policyBody string) (bool, error) {
	res, err := c.XPackIlmPutLifecycle().Policy(policyName).BodyString(policyBody).Do(ctx)
	return res.Acknowledged, err
}

func (c *Client) GetILMPolicy(ctx context.Context, policyName string) (map[string]interface{}, error) {
	res, err := c.XPackIlmGetLifecycle().Pretty(true).Human(true).Policy(policyName).Do(ctx)
	if err != nil {
		return nil, err
	}
	return res[policyName].Policy, nil
}

func (c *Client) DeleteILMPolicy(ctx context.Context, policyName string) (bool, error) {
	res, err := c.XPackIlmDeleteLifecycle().Policy(policyName).Human(true).Do(ctx)
	return res.Acknowledged, err
}
