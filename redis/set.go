// Package redis
// Date: 2023/12/4 14:01
// Author: Amu
// Description:
package redis

// ======================== set 指令 ======================== //

func (rc *Client) SAdd(key string, values ...interface{}) (int64, error) {
	return rc.UniversalClient.SAdd(ctx, key, values...).Result()
}

// SPop 随机获取一个元素，无序，随机
func (rc *Client) SPop(key string) (string, error) {
	return rc.UniversalClient.SPop(ctx, key).Result()
}

func (rc *Client) SRem(key string, values ...interface{}) (int64, error) {
	return rc.UniversalClient.SRem(ctx, key, values...).Result()
}

func (rc *Client) SMembers(key string) ([]string, error) {
	return rc.UniversalClient.SMembers(ctx, key).Result()
}

func (rc *Client) SIsMembers(key string, value interface{}) (bool, error) {
	return rc.UniversalClient.SIsMember(ctx, key, value).Result()
}

func (rc *Client) SCard(key string) (int64, error) {
	return rc.UniversalClient.SCard(ctx, key).Result()
}

func (rc *Client) SUnion(key1, key2 string) ([]string, error) {
	return rc.UniversalClient.SUnion(ctx, key1, key2).Result()
}

func (rc *Client) SDiff(key1, key2 string) ([]string, error) {
	return rc.UniversalClient.SDiff(ctx, key1, key2).Result()
}

func (rc *Client) SInter(key1, key2 string) ([]string, error) {
	return rc.UniversalClient.SInter(ctx, key1, key2).Result()
}
