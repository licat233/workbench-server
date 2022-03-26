package main

import (
	"flag"
	"fmt"
	"time"

	"workbench-server/app/api/v1/internal/config"
	"workbench-server/app/api/v1/internal/handler"
	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/common/errorx"
	"workbench-server/common/global"
	"workbench-server/common/logxs"
	"workbench-server/common/utils"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/app-api.yaml", "the config file")

func main() {
	//修改go语言默认时区为CST
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone

	flag.Parse()
	// 初始化日志，必须放在前面这个地方
	logxs.InitLogger()

	//手动设置配置文件路径
	*configFile = "../../../common/config/config.yaml"

	//初始化global
	global.Initialize(*configFile)

	//初始化config
	var c config.Config
	conf.MustLoad(*configFile, &c)

	//配置端口
	c.Port = 9100

	//初始化静态文件存储目录
	utils.PathHandler(&c.Public, "")

	//初始化ctx
	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	//注册路由
	handler.RegisterHandlers(server, ctx)

	//错误处理
	httpx.SetErrorHandler(errorx.ErrorHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
