package rbac

import (
	"context"
	"strconv"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

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

func (l *GetPolicyByRoleIdLogic) GetPolicyByRoleId(req *types.IdReq) (any, error) {
	//TODO casbin中未找到该接口
	rid := strconv.Itoa(int(req.Id))
	return respx.DefaultSingleResp(nil, l.svcCtx.Casbin.GetPermissionsForUser(rid))
}
