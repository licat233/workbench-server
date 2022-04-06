package user

import (
	"context"
	"fmt"
	"net/http"

	"workbench-server/common/errorx"
	"workbench-server/common/jwtx/model"
	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"

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

func (l *LogoutLogic) Logout(r *http.Request) (any, error) {
	err := LogoutHandler(l.ctx, r)
	if err != nil {
		return nil, errorx.InternalError("注销失败:Redis错误", err)
	}
	return respx.StateResp("注销成功"), nil
}

//LogoutHandler 注销处理
///返回 errerrorx.InternalError
func LogoutHandler(ctx context.Context, r *http.Request) error {
	jwtToken := r.Header.Get("x-token")
	uid := ctx.Value("uid")
	err := model.InsertToRedis(jwtToken, fmt.Sprint(uid), model.LOGOUT)
	if err != nil {
		return err
	}
	return nil
}
