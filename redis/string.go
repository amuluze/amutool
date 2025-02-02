// Package redis
// Date: 2023/12/4 14:00
// Author: Amu
// Description:
package redis

import "time"

// ======================== string 指令 ======================== //

func (rc *Client) Set(key, value string) (string, error) {
	return rc.UniversalClient.Set(ctx, key, value, 0).Result()
}

func (rc *Client) SetNX(key, value string, duration time.Duration) (bool, error) {
	return rc.UniversalClient.SetNX(ctx, key, value, duration).Result()
}

func (rc *Client) SetEX(key, value string, duration time.Duration) (string, error) {
	return rc.UniversalClient.SetEX(ctx, key, value, duration).Result()
}

func (rc *Client) Get(key string) (string, error) {
	return rc.UniversalClient.Get(ctx, key).Result()
}

func (rc *Client) GetRange(key string, startIndex int64, endIndex int64) (string, error) {
	return rc.UniversalClient.GetRange(ctx, key, startIndex, endIndex).Result()
}

func (rc *Client) Incr(key string) (int64, error) {
	return rc.UniversalClient.Incr(ctx, key).Result()
}

func (rc *Client) IncrBy(key string, step int64) (int64, error) {
	return rc.UniversalClient.IncrBy(ctx, key, step).Result()
}

func (rc *Client) Decr(key string) (int64, error) {
	return rc.UniversalClient.Decr(ctx, key).Result()
}

func (rc *Client) DecrBy(key string, step int64) (int64, error) {
	return rc.UniversalClient.DecrBy(ctx, key, step).Result()
}

func (rc *Client) Append(key string, appendString string) (int64, error) {
	return rc.UniversalClient.Append(ctx, key, appendString).Result()
}

func (rc *Client) StrLen(key string) (int64, error) {
	return rc.UniversalClient.StrLen(ctx, key).Result()
}
