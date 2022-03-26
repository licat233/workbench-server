package rbac

import (
	"context"

	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"
	"workbench-server/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUsersByDomainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllUsersByDomainLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllUsersByDomainLogic {
	return GetAllUsersByDomainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllUsersByDomainLogic) GetAllUsersByDomain(req *types.DomainReq) (interface{}, error) {
	return respx.NewDefaultResponse(nil, l.svcCtx.CasbinObj.GetAllUsersByDomain(req.Domain))
}
