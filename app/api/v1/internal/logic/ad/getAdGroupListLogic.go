package ad

import (
	"context"

	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"
	"workbench-server/common/errorx"
	"workbench-server/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdGroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdGroupListLogic {
	return &GetAdGroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdGroupListLogic) GetAdGroupList(req *types.ListReq) (interface{}, error) {
	data, err := l.svcCtx.AdGroupModel.FindAll(req.Current, req.PageSize)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库查询失败", err)
	}
	total, err := l.svcCtx.AdGroupModel.FindCount()
	return respx.NewListResponse(err, data, total, req.PageSize)
}
