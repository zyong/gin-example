package ginredis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// 目的是实现一个redis的连接池服务
// 1、可用组件go-redis
// 2、使用原生命令压测，压测使用benchmark

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "192.168.3.13:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
	PoolSize: 1,
})

var ctx = context.TODO()

func Set(key string, value interface{}, timeout time.Duration) {
	redisClient.Set(ctx, key, value, timeout)
}

func Get(key string) interface{} {
	ret := redisClient.Get(ctx, key)
	return ret.Val()
}
