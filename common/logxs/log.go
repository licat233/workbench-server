package logxs

import (
	"github.com/zeromicro/go-zero/core/logx"
)

func InitLogger() {
	c := logx.LogConf{
		Mode:  "file",
		Path:  "logs",
		Level: "error",
	}
	logx.MustSetup(c)
}
