package initialize

import "time"

/** Initialize 初始化服务配置
 * description: 务必按照逻辑顺序执行
 */
func Initialize() {
	//修改go语言默认时区为CST
	time.Local = time.FixedZone("CST", 8*3600) //东八区
	Config()                                   //必须先初始化
	Zap()                                      //其次
	Redis()                                    //依赖于zap
	RedisX()                                   //初始化go-zero封装的redis客户端
	GormDB()                                   //初始化gorm客户端
	Casbin()                                   //初始化casbin Enforcer，依赖GORM
	SQLConn()                                  //初始化go-zero封装的mysql客户端：sqlx.SqlConn
}
