package initialize

import (
	"workbench-server/common/global"
	"workbench-server/common/gormx"
)

func GormDB() {
	global.GORM = gormx.CreateGorm()
}
