// Package redis
// Date: 2023/12/4 14:00
// Author: Amu
// Description:
package redis

// ======================== list 指令 ======================== //

func (rc *Client) LPush(key string, values ...interface{}) (int64, error) {
	return rc.UniversalClient.LPush(ctx, key, values...).Result()
}

func (rc *Client) RPush(key string, values ...interface{}) (int64, error) {
	return rc.UniversalClient.RPush(ctx, key, values...).Result()
}

// LInsert 在第一个 target 的前或后插入新元素 value
func (rc *Client) LInsert(key string, location string, target interface{}, value interface{}) (int64, error) {
	return rc.UniversalClient.LInsert(ctx, key, location, target, value).Result()
}

func (rc *Client) LInsertBefore(key string, target interface{}, value interface{}) (int64, error) {
	return rc.UniversalClient.LInsertBefore(ctx, key, target, value).Result()
}

func (rc *Client) LInsertAfter(key string, target interface{}, value interface{}) (int64, error) {
	return rc.UniversalClient.LInsertAfter(ctx, key, target, value).Result()
}

func (rc *Client) LSet(key string, index int64, value interface{}) (string, error) {
	return rc.UniversalClient.LSet(ctx, key, index, value).Result()
}

func (rc *Client) LLen(key string) (int64, error) {
	return rc.UniversalClient.LLen(ctx, key).Result()
}

func (rc *Client) LIndex(key string, index int64) (string, error) {
	return rc.UniversalClient.LIndex(ctx, key, index).Result()
}

func (rc *Client) LRange(key string, start int64, end int64) ([]string, error) {
	return rc.UniversalClient.LRange(ctx, key, start, end).Result()
}

func (rc *Client) LPop(key string) (string, error) {
	return rc.UniversalClient.LPop(ctx, key).Result()
}

func (rc *Client) RPop(key string) (string, error) {
	return rc.UniversalClient.RPop(ctx, key).Result()
}

// LRem 删除指定数量 nums 的 value，返回实际删除的元素个数
func (rc *Client) LRem(key string, nums int64, value string) (int64, error) {
	return rc.UniversalClient.LRem(ctx, key, nums, value).Result()
}
