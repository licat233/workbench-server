package ad

import (
	"context"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

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

func (l *GetAdAccountListLogic) GetAdAccountList(req *types.ListReq) (any, error) {
	data, err := l.svcCtx.AdAccountModel.FindAll(req.Current, req.PageSize)
	if err != nil {
		return nil, errorx.InternalError("数据库查询失败", err)
	}
	total, err := l.svcCtx.AdAccountModel.FindCount()
	return respx.DefaultListResp(err, data, total, req.PageSize, req.Current)
}
