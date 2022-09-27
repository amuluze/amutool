// Package redis
// Date: 2022/9/26 17:51
// Author: Amu
// Description:
package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
)

var rc redis.UniversalClient

// ======================== 基本指令 ======================== //

func Keys() ([]string, error) {
	return rc.Keys(ctx, "*").Result()
}

func Type(key string) (string, error) {
	return rc.Type(ctx, key).Result()
}

func Delete(keys ...string) (int64, error) {
	return rc.Del(ctx, keys...).Result()
}

func Exists(key string) (int64, error) {
	return rc.Exists(ctx, key).Result()
}

func Expire(key string, expireDuration time.Duration) (bool, error) {
	return rc.Expire(ctx, key, expireDuration).Result()
}

func ExpireAt(key string, expireTime time.Time) (bool, error) {
	return rc.ExpireAt(ctx, key, expireTime).Result()
}

func TTL(key string) (time.Duration, error) {
	return rc.TTL(ctx, key).Result()
}

func PTTL(key string) (time.Duration, error) {
	return rc.PTTL(ctx, key).Result()
}

func DBSize() (int64, error) {
	return rc.DBSize(ctx).Result()
}

// FlushDB 清空当前数据库
func FlushDB() (string, error) {
	return rc.FlushDB(ctx).Result()
}

// FlushAll 清空所有数据库
func FlushAll() (string, error) {
	return rc.FlushAll(ctx).Result()
}

// ======================== string 指令 ======================== //

func Set(key, value string) (string, error) {
	return rc.Set(ctx, key, value, 0).Result()
}

func SetNX(key, value string, duration time.Duration) (bool, error) {
	return rc.SetNX(ctx, key, value, duration).Result()
}

func SetEX(key, value string, duration time.Duration) (string, error) {
	return rc.SetEX(ctx, key, value, duration).Result()
}

func Get(key string) (string, error) {
	return rc.Get(ctx, key).Result()
}

func GetRange(key string, startIndex int64, endIndex int64) (string, error) {
	return rc.GetRange(ctx, key, startIndex, endIndex).Result()
}

func Incr(key string) (int64, error) {
	return rc.Incr(ctx, key).Result()
}

func IncrBy(key string, step int64) (int64, error) {
	return rc.IncrBy(ctx, key, step).Result()
}

func Decr(key string) (int64, error) {
	return rc.Decr(ctx, key).Result()
}

func DecrBy(key string, step int64) (int64, error) {
	return rc.DecrBy(ctx, key, step).Result()
}

func Append(key string, appendString string) (int64, error) {
	return rc.Append(ctx, key, appendString).Result()
}

func StrLen(key string) (int64, error) {
	return rc.StrLen(ctx, key).Result()
}

// ======================== list 指令 ======================== //

func LPush(key string, values ...interface{}) (int64, error) {
	return rc.LPush(ctx, key, values...).Result()
}

func RPush(key string, values ...interface{}) (int64, error) {
	return rc.RPush(ctx, key, values...).Result()
}

// LInsert 在第一个 target 的前或后插入新元素 value
func LInsert(key string, location string, target interface{}, value interface{}) (int64, error) {
	return rc.LInsert(ctx, key, location, target, value).Result()
}

func LInsertBefore(key string, target interface{}, value interface{}) (int64, error) {
	return rc.LInsertBefore(ctx, key, target, value).Result()
}

func LInsertAfter(key string, target interface{}, value interface{}) (int64, error) {
	return rc.LInsertAfter(ctx, key, target, value).Result()
}

func LSet(key string, index int64, value interface{}) (string, error) {
	return rc.LSet(ctx, key, index, value).Result()
}

func LLen(key string) (int64, error) {
	return rc.LLen(ctx, key).Result()
}

func LIndex(key string, index int64) (string, error) {
	return rc.LIndex(ctx, key, index).Result()
}

func LRange(key string, start int64, end int64) ([]string, error) {
	return rc.LRange(ctx, key, start, end).Result()
}

func LPop(key string) (string, error) {
	return rc.LPop(ctx, key).Result()
}

func RPop(key string) (string, error) {
	return rc.RPop(ctx, key).Result()
}

// LRem 删除指定数量 nums 的 value，返回实际删除的元素个数
func LRem(key string, nums int64, value string) (int64, error) {
	return rc.LRem(ctx, key, nums, value).Result()
}

// ======================== set 指令 ======================== //

func SAdd(key string, values ...interface{}) (int64, error) {
	return rc.SAdd(ctx, key, values...).Result()
}

// SPop 随机获取一个元素，无序，随机
func SPop(key string) (string, error) {
	return rc.SPop(ctx, key).Result()
}

func SRem(key string, values ...interface{}) (int64, error) {
	return rc.SRem(ctx, key, values...).Result()
}

func SMembers(key string) ([]string, error) {
	return rc.SMembers(ctx, key).Result()
}

func SIsMembers(key string, value interface{}) (bool, error) {
	return rc.SIsMember(ctx, key, value).Result()
}

func SCard(key string) (int64, error) {
	return rc.SCard(ctx, key).Result()
}

func SUnion(key1, key2 string) ([]string, error) {
	return rc.SUnion(ctx, key1, key2).Result()
}

func SDiff(key1, key2 string) ([]string, error) {
	return rc.SDiff(ctx, key1, key2).Result()
}

func SInter(key1, key2 string) ([]string, error) {
	return rc.SInter(ctx, key1, key2).Result()
}

// ======================== zset 指令 ======================== //

func ZAdd(key string, member interface{}, score float64) (int64, error) {
	return rc.ZAdd(ctx, key, &redis.Z{Score: score, Member: member}).Result()
}

// ZIncrBy 增加 member 的分值，返回更新后的分值
func ZIncrBy(key string, member string, score float64) (float64, error) {
	return rc.ZIncrBy(ctx, key, score, member).Result()
}

// ZRange 获取根据 score 排序后的 [start, end] 元素
func ZRange(key string, start int64, end int64) ([]string, error) {
	return rc.ZRange(ctx, key, start, end).Result()
}

func ZRevRange(key string, start int64, end int64) ([]string, error) {
	return rc.ZRange(ctx, key, start, end).Result()
}

func ZRangeByScore(key string, minScore string, maxScore string) ([]string, error) {
	return rc.ZRangeByScore(ctx, key, &redis.ZRangeBy{Min: minScore, Max: maxScore}).Result()
}

func ZRevRangeByScore(key string, minScore string, maxScore string) ([]string, error) {
	return rc.ZRangeByScore(ctx, key, &redis.ZRangeBy{Min: minScore, Max: maxScore}).Result()
}

func ZCard(key string) (int64, error) {
	return rc.ZCard(ctx, key).Result()
}

func ZCount(key string, minScore, maxScore string) (int64, error) {
	return rc.ZCount(ctx, key, minScore, maxScore).Result()
}

func ZScore(key string, score string) (float64, error) {
	return rc.ZScore(ctx, key, score).Result()
}

func ZRank(key string, value string) (int64, error) {
	return rc.ZRank(ctx, key, value).Result()
}

func ZRevRank(key string, value string) (int64, error) {
	return rc.ZRank(ctx, key, value).Result()
}

func ZRem(key string, value string) (int64, error) {
	return rc.ZRem(ctx, key, value).Result()
}

func ZRemRangeByRank(key string, startIndex, endIndex int64) (int64, error) {
	return rc.ZRemRangeByRank(ctx, key, startIndex, endIndex).Result()
}

func ZRemRangeByScore(key string, minScore, maxScore string) (int64, error) {
	return rc.ZRemRangeByScore(ctx, key, minScore, maxScore).Result()
}

// ======================== hash 指令 ======================== //

func Hset(key string, field1, value1 string, field2, value2 string) (int64, error) {
	return rc.HSet(ctx, key, field1, value1, field2, value2).Result()
}

func HMset(key string, values map[string]interface{}) (bool, error) {
	return rc.HMSet(ctx, key, values).Result()
}

func HGet(key, field string) (interface{}, error) {
	return rc.HGet(ctx, key, field).Result()
}

func HGetAll(key string) (map[string]string, error) {
	return rc.HGetAll(ctx, key).Result()
}

func HDel(key string, field ...string) (int64, error) {
	return rc.HDel(ctx, key, field...).Result()
}

func HExists(key, field string) (bool, error) {
	return rc.HExists(ctx, key, field).Result()
}

func HLen(key string) (int64, error) {
	return rc.HLen(ctx, key).Result()
}
