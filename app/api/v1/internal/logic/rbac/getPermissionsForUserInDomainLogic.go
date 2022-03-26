package rbac

import (
	"context"

	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"
	"workbench-server/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermissionsForUserInDomainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPermissionsForUserInDomainLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetPermissionsForUserInDomainLogic {
	return GetPermissionsForUserInDomainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPermissionsForUserInDomainLogic) GetPermissionsForUserInDomain(req *types.RoleAndDomainReq) (interface{}, error) {
	return respx.NewDefaultResponse(nil, l.svcCtx.CasbinObj.GetPermissionsForUserInDomain(req.Sub, req.Dom))
}
