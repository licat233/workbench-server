package ad

import (
	"context"

	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"
	"workbench-server/common/errorx"
	"workbench-server/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdAccountListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdAccountListLogic {
	return &GetAdAccountListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdAccountListLogic) GetAdAccountList(req *types.ListReq) (interface{}, error) {
	data, err := l.svcCtx.AdAccountModel.FindAll(req.Current, req.PageSize)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库查询失败", err)
	}
	total, err := l.svcCtx.AdAccountModel.FindCount()
	return respx.NewListResponse(err, data, total, req.PageSize)
}
