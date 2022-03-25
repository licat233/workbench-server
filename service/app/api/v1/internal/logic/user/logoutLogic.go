package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"workbench-server/global/errorx"
	"workbench-server/global/respx"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) LogoutLogic {
	return LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(r *http.Request) (interface{}, error) {
	exp, err := strconv.ParseInt(fmt.Sprintf("%v", l.ctx.Value("ExpiresAt")), 10, 64)
	if err != nil {
		return nil, errorx.NewDefaultError("获取过期时间失败", err)
	}
	now := time.Now().Unix()
	dura := exp - now
	if dura <= 0 {
		return nil, errorx.NewDefaultError("当前登录已过期，无需主动注销", nil)
	}
	jwtToken := r.Header.Get("Authorization")
	token := fmt.Sprintf("token:blacklist:%s", jwtToken)
	err = l.svcCtx.RedisClient.Setex(token, "logout", int(dura))
	return respx.NewDefaultResponse(err, nil, "注销成功", "注销失败:Redis错误")
}
