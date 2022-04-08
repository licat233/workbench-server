package redisx

import (
	"context"
	"workbench-server/common/global"

	"github.com/go-redis/redis/v8"
	redisx "github.com/zeromicro/go-zero/core/stores/redis"
	"go.uber.org/zap"
)

//CreateRedis 创建redis客户端
func CreateRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.CONFIG.Redis.Addr,
		Password: global.CONFIG.Redis.Password,
		DB:       global.CONFIG.Redis.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.ZAP.Error("redis connect ping failed, err:", zap.Error(err))
		panic("连接redis服务器失败")
	}
	return client
}

//CreateRedisX 创建go-zero封装的redis客户端
func CreateRedisX() *redisx.Redis {
	return redisx.New(global.CONFIG.Redis.Addr, redisx.WithPass(global.CONFIG.Redis.Password))
}
