package user

import (
	"context"
	"errors"
	"net/http"

	"workbench-server/common/errorx"
	"workbench-server/common/jwtx"
	"workbench-server/common/jwtx/model"
	"workbench-server/common/respx"
	"workbench-server/common/tool"
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
	uid, err := tool.GetJwtUid(ctx)
	if err != nil {
		return err
	}
	claims, err := jwtx.NewJWT().ParseToken(jwtToken)
	if err != nil {
		return err
	}
	exp := jwtx.ExpiredTime(*claims)
	if exp == 0 {
		return errors.New("获取jwt过期时间失败")
	}
	err = model.InsertToRedis(uid, jwtToken, exp, model.LOGOUT)
	if err != nil {
		return err
	}
	return nil
}
