package initialize

import (
	"workbench-server/common/casbinx"
	"workbench-server/common/global"
)

func Casbin() {
	global.CASBIN = casbinx.Casbin()
}
