package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf // rest api配置
	// Auth          config.JWT
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	// Mysql         struct {
	// 	DataSource string
	// }
	// Redis cache.CacheConf // redis缓存
}
