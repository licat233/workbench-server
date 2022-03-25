package rbac

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *RemovePolicyLogic) RemovePolicy(req types.PolicyReq) (interface{}, error) {
	ok, err := l.svcCtx.CasbinObj.RemovePolicy(req.Sub, req.Obj, req.Act)
	return respx.NewDefaultResponse(err, ok)
}
