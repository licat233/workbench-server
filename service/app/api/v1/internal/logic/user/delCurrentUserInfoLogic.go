package user

import (
	"context"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/common/utils"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type DelCurrentUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelCurrentUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) DelCurrentUserInfoLogic {
	return DelCurrentUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelCurrentUserInfoLogic) DelCurrentUserInfo() (interface{}, error) {
	uid, err := utils.GetJwtUid(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError("获取当前用户ID失败", err)
	}
	err = l.svcCtx.UserModel.Delete(uid)
	return respx.NewDefaultResponse(err, nil, "删除成功", "删除失败")
}
