package rbac

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllActionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllActionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllActionsLogic {
	return GetAllActionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllActionsLogic) GetAllActions() (any, error) {
	data := l.svcCtx.Casbin.GetAllActions()
	return respx.DefaultSingleResp(nil, data)
}
