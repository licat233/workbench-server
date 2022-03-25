package rbac

import (
	"context"

	"workbench-server/global/respx"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddPolicyLogic {
	return AddPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPolicyLogic) AddPolicy(req types.PolicyReq) (interface{}, error) {
	ok, err := l.svcCtx.CasbinObj.AddPolicy(req.Sub, req.Dom, req.Obj, req.Act)
	return respx.NewDefaultResponse(err, ok)
}
