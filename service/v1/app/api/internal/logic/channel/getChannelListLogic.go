package channel

import (
	"context"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChannelListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChannelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetChannelListLogic {
	return GetChannelListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChannelListLogic) GetChannelList(req *types.ListReq) (any, error) {
	data, err := l.svcCtx.AdChannelModel.FindAll(req.Current, req.PageSize)
	if err != nil {
		return nil, errorx.InternalError("数据库查询失败", err)
	}
	total, err := l.svcCtx.AdChannelModel.FindCount()
	return respx.DefaultListResp(err, data, total, req.PageSize, req.Current)
}
