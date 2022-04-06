package ad

import (
	"context"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

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

func (l *GetAdCampaignListLogic) GetAdCampaignList(req *types.ListReq) (any, error) {
	data, err := l.svcCtx.AdCampaignModel.FindAll(req.Current, req.PageSize)
	if err != nil {
		return nil, errorx.InternalError("数据库查询失败", err)
	}
	total, err := l.svcCtx.AdCampaignModel.FindCount()
	return respx.DefaultListResp(err, data, total, req.PageSize, req.Current)
}
