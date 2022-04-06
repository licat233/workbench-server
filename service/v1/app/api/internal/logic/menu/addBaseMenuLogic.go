package menu

import (
	"context"

	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBaseMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddBaseMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddBaseMenuLogic {
	return AddBaseMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddBaseMenuLogic) AddBaseMenu() (*types.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &types.BaseResp{}, nil
}
