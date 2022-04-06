package initialize

import (
	"workbench-server/common/global"
	"workbench-server/common/zapx"
)

func Zap() {
	global.ZAP = zapx.Zap()
}
