package logxs

import (
	"github.com/tal-tech/go-zero/core/logx"
)

func InitLogger() {
	c := logx.LogConf{
		Mode:  "file",
		Path:  "logs",
		Level: "error",
	}
	logx.MustSetup(c)
}
