package menu

import (
	"context"

	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddMenuAuthorityLogic {
	return AddMenuAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMenuAuthorityLogic) AddMenuAuthority() (*types.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &types.BaseResp{}, nil
}
