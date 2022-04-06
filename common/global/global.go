package global

import (
	"sync"
	"workbench-server/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	redisx "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	CONFIG *config.Server //全局配置，务必先初始化
	VIPER  *viper.Viper   //全局配置器
	ZAP    *zap.Logger    //全局日志
	GORM   *gorm.DB       //全局mysql orm
	REDIS  *redis.Client  //全局redis

	DBList map[string]*gorm.DB //gorm数据库列表

	SQLConn sqlx.SqlConn  //gozero定义后的mysql
	RedisX  *redisx.Redis //gozero定义后的redis

	Concurrency_Control = &singleflight.Group{} //并发控制
	lock                sync.RWMutex            //全局读写锁
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
