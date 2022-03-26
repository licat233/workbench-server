package rbac

import (
	"context"

	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetPolicyLogic {
	return GetPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPolicyLogic) GetPolicy() (interface{}, error) {
	return respx.NewDefaultResponse(nil, l.svcCtx.CasbinObj.GetPolicy())
}
