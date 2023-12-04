// Package redis
// Date: 2023/12/4 13:59
// Author: Amu
// Description:
package redis

import "time"

// ======================== 基本指令 ======================== //

func (rc *Client) Keys() ([]string, error) {
	return rc.UniversalClient.Keys(ctx, "*").Result()
}

func (rc *Client) Type(key string) (string, error) {
	return rc.UniversalClient.Type(ctx, key).Result()
}

func (rc *Client) Delete(keys ...string) (int64, error) {
	return rc.UniversalClient.Del(ctx, keys...).Result()
}

func (rc *Client) Exists(key string) (int64, error) {
	return rc.UniversalClient.Exists(ctx, key).Result()
}

func (rc *Client) Expire(key string, expireDuration time.Duration) (bool, error) {
	return rc.UniversalClient.Expire(ctx, key, expireDuration).Result()
}

func (rc *Client) ExpireAt(key string, expireTime time.Time) (bool, error) {
	return rc.UniversalClient.ExpireAt(ctx, key, expireTime).Result()
}

func (rc *Client) TTL(key string) (time.Duration, error) {
	return rc.UniversalClient.TTL(ctx, key).Result()
}

func (rc *Client) PTTL(key string) (time.Duration, error) {
	return rc.UniversalClient.PTTL(ctx, key).Result()
}

func (rc *Client) DBSize() (int64, error) {
	return rc.UniversalClient.DBSize(ctx).Result()
}

// FlushDB 清空当前数据库
func (rc *Client) FlushDB() (string, error) {
	return rc.UniversalClient.FlushDB(ctx).Result()
}

// FlushAll 清空所有数据库
func (rc *Client) FlushAll() (string, error) {
	return rc.UniversalClient.FlushAll(ctx).Result()
}
