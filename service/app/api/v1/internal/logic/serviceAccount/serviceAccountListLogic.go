package serviceAccount

import (
	"context"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceAccountListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServiceAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ServiceAccountListLogic {
	return ServiceAccountListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServiceAccountListLogic) ServiceAccountList(req types.ListReq) (interface{}, error) {
	data, err := l.svcCtx.ServiceAccountModel.FindAll(req.Current, req.PageSize)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库查询失败", err)
	}
	total, err := l.svcCtx.ServiceAccountModel.FindCount()
	return respx.NewListResponse(err, data, total, req.PageSize)
}
