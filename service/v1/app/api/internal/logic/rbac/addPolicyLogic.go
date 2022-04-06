package rbac

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

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

func (l *AddPolicyLogic) AddPolicy(req *types.PolicyReq) (any, error) {
	ok, err := l.svcCtx.Casbin.AddPolicy(req.Sub, req.Dom, req.Obj, req.Act)
	return respx.DefaultSingleResp(err, ok)
}
