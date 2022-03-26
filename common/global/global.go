package global

import (
	"workbench-server/common/config"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.GlobalConfig
	GormDB *gorm.DB
	Redis  *redis.Client
	ZapLog *zap.Logger
	Casbin *casbin.Enforcer
)

func Initialize(path string) {
	Config = NewGlobalConfig(path)
	Casbin = NewCasbin()
}

func NewGlobalConfig(path string) *config.GlobalConfig {
	var c = new(config.GlobalConfig)
	conf.MustLoad(path, c)
	return c
}

func NewCasbin() *casbin.Enforcer {
	a, err := gormadapter.NewAdapter("mysql", Config.Mysql.DataSource, true)
	if err != nil {
		logx.Errorf("NewCasbin连接数据库错误: %v", err)
		panic("NewCasbin连接数据库错误")
		// return nil
	}

	CasbinObj, err := casbin.NewEnforcer(Config.Casbin.ModelConf, a)
	if err != nil {
		logx.Errorf("初始化casbin错误: %v", err)
		return nil
	}
	err = CasbinObj.LoadPolicy()
	if err != nil {
		logx.Errorf("CasbinObj.LoadPolicy()错误: %v", err)
		return nil
	}
	return CasbinObj
}
