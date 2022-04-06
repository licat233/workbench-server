package rbac

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

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

func (l *GetPermissionsForUserInDomainLogic) GetPermissionsForUserInDomain(req *types.RoleAndDomainReq) (any, error) {
	return respx.DefaultSingleResp(nil, l.svcCtx.Casbin.GetPermissionsForUserInDomain(req.Sub, req.Dom))
}
