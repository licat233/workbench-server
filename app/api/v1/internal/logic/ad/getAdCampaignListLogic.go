package ad

import (
	"context"

	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"
	"workbench-server/common/errorx"
	"workbench-server/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdCampaignListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdCampaignListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdCampaignListLogic {
	return &GetAdCampaignListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdCampaignListLogic) GetAdCampaignList(req *types.ListReq) (interface{}, error) {
	data, err := l.svcCtx.AdCampaignModel.FindAll(req.Current, req.PageSize)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库查询失败", err)
	}
	total, err := l.svcCtx.AdCampaignModel.FindCount()
	return respx.NewListResponse(err, data, total, req.PageSize)
}
