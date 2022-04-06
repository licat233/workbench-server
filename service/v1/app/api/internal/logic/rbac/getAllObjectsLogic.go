package rbac

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"

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

func (l *GetAllObjectsLogic) GetAllObjects() (any, error) {
	return respx.DefaultSingleResp(nil, l.svcCtx.Casbin.GetAllObjects())
}
