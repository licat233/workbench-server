package main

import (
	"fmt"
	"testing"
	"workbench-server/common/errorx"
)

func Test_errorx(t *testing.T) {
	err := &errorx.Errorx{
		Success:      false,
		ErrorCode:    400,
		ErrorMessage: "请求失败",
		TraceMessage: "账号或者密码错误",
		ShowType:     2,
		TraceId:      "",
		Host:         "",
	}
	fmt.Println(err.Data())
}
