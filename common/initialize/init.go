package initialize

import "time"

func Initialize() {
	//修改go语言默认时区为CST
	time.Local = time.FixedZone("CST", 8*3600) // 东八
	Config()
	Redis()
	RedisX()
	GormDB()
	Zap()
	SQLConn()
}
