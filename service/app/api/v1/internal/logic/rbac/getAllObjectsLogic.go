package rbac

import (
	"context"

	"workbench-server/global/respx"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllObjectsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllObjectsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllObjectsLogic {
	return GetAllObjectsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllObjectsLogic) GetAllObjects() (interface{}, error) {
	return respx.NewDefaultResponse(nil, l.svcCtx.CasbinObj.GetAllObjects())
}
