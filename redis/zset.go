// Package redis
// Date: 2023/12/4 14:01
// Author: Amu
// Description:
package redis

import "github.com/go-redis/redis/v8"

// ======================== zset 指令 ======================== //

func (rc *Client) ZAdd(key string, member interface{}, score float64) (int64, error) {
	return rc.UniversalClient.ZAdd(ctx, key, &redis.Z{Score: score, Member: member}).Result()
}

// ZIncrBy 增加 member 的分值，返回更新后的分值
func (rc *Client) ZIncrBy(key string, member string, score float64) (float64, error) {
	return rc.UniversalClient.ZIncrBy(ctx, key, score, member).Result()
}

// ZRange 获取根据 score 排序后的 [start, end] 元素
func (rc *Client) ZRange(key string, start int64, end int64) ([]string, error) {
	return rc.UniversalClient.ZRange(ctx, key, start, end).Result()
}

func (rc *Client) ZRevRange(key string, start int64, end int64) ([]string, error) {
	return rc.UniversalClient.ZRange(ctx, key, start, end).Result()
}

func (rc *Client) ZRangeByScore(key string, minScore string, maxScore string) ([]string, error) {
	return rc.UniversalClient.ZRangeByScore(ctx, key, &redis.ZRangeBy{Min: minScore, Max: maxScore}).Result()
}

func (rc *Client) ZRevRangeByScore(key string, minScore string, maxScore string) ([]string, error) {
	return rc.UniversalClient.ZRangeByScore(ctx, key, &redis.ZRangeBy{Min: minScore, Max: maxScore}).Result()
}

func (rc *Client) ZCard(key string) (int64, error) {
	return rc.UniversalClient.ZCard(ctx, key).Result()
}

func (rc *Client) ZCount(key string, minScore, maxScore string) (int64, error) {
	return rc.UniversalClient.ZCount(ctx, key, minScore, maxScore).Result()
}

func (rc *Client) ZScore(key string, score string) (float64, error) {
	return rc.UniversalClient.ZScore(ctx, key, score).Result()
}

func (rc *Client) ZRank(key string, value string) (int64, error) {
	return rc.UniversalClient.ZRank(ctx, key, value).Result()
}

func (rc *Client) ZRevRank(key string, value string) (int64, error) {
	return rc.UniversalClient.ZRank(ctx, key, value).Result()
}

func (rc *Client) ZRem(key string, value string) (int64, error) {
	return rc.UniversalClient.ZRem(ctx, key, value).Result()
}

func (rc *Client) ZRemRangeByRank(key string, startIndex, endIndex int64) (int64, error) {
	return rc.UniversalClient.ZRemRangeByRank(ctx, key, startIndex, endIndex).Result()
}

func (rc *Client) ZRemRangeByScore(key string, minScore, maxScore string) (int64, error) {
	return rc.UniversalClient.ZRemRangeByScore(ctx, key, minScore, maxScore).Result()
}
