package initialize

import (
	"workbench-server/common/global"
	"workbench-server/common/gormx"

	"gorm.io/gorm"
)

func GormDB() {
	global.GORM = new(gorm.DB)
	gormx.GormMysql()
}
