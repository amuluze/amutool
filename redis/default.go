// Package redis
// Date: 2022/9/26 17:51
// Author: Amu
// Description:
package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
)

func (rc *Client) Close() {
	rc.Close()
}

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
