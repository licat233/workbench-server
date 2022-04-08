package main

import (
	"flag"
	"fmt"
	"net/http"

	"workbench-server/common/errorx"
	"workbench-server/common/global"
	"workbench-server/common/initialize"
	"workbench-server/service/v1/app/api/internal/config"
	"workbench-server/service/v1/app/api/internal/handler"
	"workbench-server/service/v1/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func initLogger() {
	c := logx.LogConf{
		ServiceName:         "woekbench-server",
		Mode:                "file",
		Encoding:            "plain",
		TimeFormat:          "",
		Path:                "logxs",
		Level:               "error",
		Compress:            true,
		KeepDays:            3,
		StackCooldownMillis: 200,
	}
	logx.MustSetup(c)
}

var configFile = flag.String("f", "etc/app-api.yaml", "the config file")

func main() {
	flag.Parse()

	//初始化项目所需的所有配置
	initialize.Initialize()

	// 初始化logx日志，必须放在前面这个地方
	initLogger()

	//初始化config
	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.Auth.AccessExpire = global.CONFIG.JWT.AccessExpire
	c.Auth.AccessSecret = global.CONFIG.JWT.AccessSecret

	//初始化ServiceContext
	ctx := svc.NewServiceContext(c)
	//初始化server
	server := rest.MustNewServer(c.RestConf)

	//server添加cors处理中间件
	//TODO：该CORS有问题
	// server.Use(middleware.CorsHandler)

	defer server.Stop()

	//注册路由
	handler.RegisterHandlers(server, ctx)

	//错误处理
	httpx.SetErrorHandler(errorx.ErrorToErrorX)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func RespHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
