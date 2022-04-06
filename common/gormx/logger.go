package gormx

import (
	"fmt"
	"workbench-server/common/global"

	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志，实现 Writer 接口
func (w *writer) Printf(message string, data ...interface{}) {
	if global.CONFIG.Mysql.LogZap {
		global.ZAP.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
