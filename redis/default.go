// Package redis
// Date: 2022/9/26 17:51
// Author: Amu
// Description:
package redis

import (
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var once sync.Once
var rc redis.UniversalClient

func init() {
	once.Do(func() {
		rc = NewClient(&Config{
			Addr:     []string{"127.0.0.1:6379"},
			Password: "Be1s.Az3",
			DB:       0,
		})
	})
}

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

// ======================== set 指令 ======================== //

// ======================== zset 指令 ======================== //

// ======================== hash 指令 ======================== //
