package initialize

import (
	"workbench-server/common/global"
	"workbench-server/common/tool"
	"workbench-server/config"

	"github.com/zeromicro/go-zero/core/conf"
)

func Config() {
	global.CONFIG = new(config.Server)
	//读取yaml配置文件，映射结构体
	conf.MustLoad(config.ProjectConfigFile, global.CONFIG)
	global.CONFIG.ServerConfig.RootPath = tool.CurrentProjectRootPath()
	//初始化静态文件存储目录
	tool.StaticServerPathHandler(global.CONFIG.StaticConfig.Root, global.CONFIG.ServerConfig.RootPath)
}
