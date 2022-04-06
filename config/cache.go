package config

import "github.com/zeromicro/go-zero/core/stores/cache"

type Cache struct {
	Redis cache.CacheConf `mapstructure:"Redis" json:"Redis" yaml:"Redis"` // redis缓存配置
}
