package rbac

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"

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

func (l *GetPolicyLogic) GetPolicy() (any, error) {
	return respx.DefaultSingleResp(nil, l.svcCtx.Casbin.GetPolicy())
}
