package user

import (
	"context"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/common/tool"
	"workbench-server/service/v1/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCurrentUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCurrentUserInfoLogic {
	return GetCurrentUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentUserInfoLogic) GetCurrentUserInfo() (any, error) {
	uid, err := tool.GetJwtUid(l.ctx)
	if err != nil {
		return nil, errorx.InternalError("获取当前用户ID失败", err)
	}
	data, err := l.svcCtx.UserModel.FindOneById(uid)
	// fmt.Println(data.UserCreateTime)
	return respx.DefaultSingleResp(err, data)
}
