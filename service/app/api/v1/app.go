package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"workbench-server/common/errorx"
	"workbench-server/common/logxs"
	"workbench-server/common/utils"
	"workbench-server/service/app/api/v1/internal/config"
	"workbench-server/service/app/api/v1/internal/handler"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/app-api.yaml", "the config file")

func main() {
	//修改go语言默认时区为CST
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone

	flag.Parse()
	// 初始化日志，必须放在前面这个地方
	logxs.InitLogger()
	var c config.Config
	var path string = "."
	var err error
	// 不是从命令行启动的
	if len(os.Args) <= 1 {
		if path, err = os.Getwd(); err == nil {
			// fmt.Println("当前路径:", path)
			configFilePath := path + "/etc/app-api.yaml"
			conf.MustLoad(configFilePath, &c)
		}
	} else {
		conf.MustLoad(*configFile, &c)
	}
	//初始化静态文件存储目录
	utils.PathHandler(&c.Storage, "")

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
