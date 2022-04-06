package rbac

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemovePolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemovePolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemovePolicyLogic {
	return RemovePolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemovePolicyLogic) RemovePolicy(req *types.PolicyReq) (any, error) {
	ok, err := l.svcCtx.Casbin.RemovePolicy(req.Sub, req.Obj, req.Act)
	return respx.DefaultSingleResp(err, ok)
}
