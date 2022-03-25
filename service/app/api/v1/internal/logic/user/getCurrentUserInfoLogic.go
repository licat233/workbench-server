package user

import (
	"context"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/common/utils"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
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
