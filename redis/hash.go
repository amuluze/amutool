// Package redis
// Date: 2023/12/4 14:01
// Author: Amu
// Description:
package redis

// ======================== hash 指令 ======================== //

func (rc *Client) Hset(key string, field1, value1 string, field2, value2 string) (int64, error) {
	return rc.UniversalClient.HSet(ctx, key, field1, value1, field2, value2).Result()
}

func (rc *Client) HMset(key string, values map[string]interface{}) (bool, error) {
	return rc.UniversalClient.HMSet(ctx, key, values).Result()
}

func (rc *Client) HGet(key, field string) (interface{}, error) {
	return rc.UniversalClient.HGet(ctx, key, field).Result()
}

func (rc *Client) HGetAll(key string) (map[string]string, error) {
	return rc.UniversalClient.HGetAll(ctx, key).Result()
}

func (rc *Client) HDel(key string, field ...string) (int64, error) {
	return rc.UniversalClient.HDel(ctx, key, field...).Result()
}

func (rc *Client) HExists(key, field string) (bool, error) {
	return rc.UniversalClient.HExists(ctx, key, field).Result()
}

func (rc *Client) HLen(key string) (int64, error) {
	return rc.UniversalClient.HLen(ctx, key).Result()
}
