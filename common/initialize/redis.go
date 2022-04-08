package initialize

import (
	"workbench-server/common/global"
	"workbench-server/common/redisx"
)

func Redis() {
	global.REDIS = redisx.CreateRedis()
}

func RedisX() {
	global.RedisX = redisx.CreateRedisX()
}
