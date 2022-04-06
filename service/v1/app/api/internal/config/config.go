package config

import (
	"workbench-server/config"

	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf // rest api配置
	Auth          config.JWT
	// Mysql         struct {
	// 	DataSource string
	// }
	// Redis cache.CacheConf // redis缓存
}
