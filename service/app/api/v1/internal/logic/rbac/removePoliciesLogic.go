package rbac

import (
	"context"

	"workbench-server/global/respx"
	"workbench-server/global/utils"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemovePoliciesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemovePoliciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemovePoliciesLogic {
	return RemovePoliciesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemovePoliciesLogic) RemovePolicies(req types.PoliciesReq) (interface{}, error) {
	rules := make([][]string, 0)
	for _, v := range req.Rules {
		rules = append(rules, utils.StrctToSlice(v))
	}
	ok, err := l.svcCtx.CasbinObj.RemovePolicies(rules)
	return respx.NewDefaultResponse(err, ok)
}
