package rbac

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *GetAllActionsLogic) GetAllActions() (interface{}, error) {
	data := l.svcCtx.CasbinObj.GetAllActions()
	return respx.NewDefaultResponse(nil, data)
}
