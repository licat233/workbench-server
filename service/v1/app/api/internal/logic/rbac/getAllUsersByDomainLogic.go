package rbac

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

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

func (l *GetAllUsersByDomainLogic) GetAllUsersByDomain(req *types.DomainReq) (any, error) {
	return respx.DefaultSingleResp(nil, l.svcCtx.Casbin.GetAllUsersByDomain(req.Domain))
}
