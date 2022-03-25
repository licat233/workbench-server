package rbac

import (
	"context"
	"strconv"

	"workbench-server/global/respx"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPolicyByRoleIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPolicyByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetPolicyByRoleIdLogic {
	return GetPolicyByRoleIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPolicyByRoleIdLogic) GetPolicyByRoleId(req types.IdReq) (interface{}, error) {
	//TODO casbin中未找到该接口
	rid := strconv.Itoa(int(req.Id))
	return respx.NewDefaultResponse(nil, l.svcCtx.CasbinObj.GetPermissionsForUser(rid))
}
