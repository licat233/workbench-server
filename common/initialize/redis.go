package initialize

import (
	"context"
	"workbench-server/common/global"

	"github.com/go-redis/redis/v8"
	redisx "github.com/zeromicro/go-zero/core/stores/redis"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.ZAP.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.ZAP.Info("redis connect ping response:", zap.String("pong", pong))
		global.REDIS = client
	}
}

func RedisX() {
	global.RedisX = redisx.New(global.CONFIG.Redis.Addr, redisx.WithPass(global.CONFIG.Redis.Password))
}
