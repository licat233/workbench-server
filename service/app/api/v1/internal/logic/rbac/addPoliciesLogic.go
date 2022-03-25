package rbac

import (
	"context"

	"workbench-server/global/respx"
	"workbench-server/global/utils"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPoliciesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPoliciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddPoliciesLogic {
	return AddPoliciesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPoliciesLogic) AddPolicies(req types.PoliciesReq) (interface{}, error) {
	// todo: add your logic here and delete this line
	rules := make([][]string, 0)
	for _, v := range req.Rules {
		rules = append(rules, utils.StrctToSlice(v))
	}
	ok, err := l.svcCtx.CasbinObj.AddPolicies(rules)
	return respx.NewDefaultResponse(err, ok)
}
