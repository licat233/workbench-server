package user

import (
	"context"

	"workbench-server/global/errorx"
	"workbench-server/global/respx"
	"workbench-server/global/utils"
	"workbench-server/service/app/api/v1/internal/svc"

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

func (l *GetCurrentUserInfoLogic) GetCurrentUserInfo() (interface{}, error) {
	uid, err := utils.GetJwtUid(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError("获取当前用户ID失败", err)
	}
	data, err := l.svcCtx.UserModel.FindOneById(uid)
	// fmt.Println(data.UserCreateTime)
	return respx.NewDefaultResponse(err, data)
}
