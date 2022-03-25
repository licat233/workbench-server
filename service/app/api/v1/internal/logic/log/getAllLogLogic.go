package log

import (
	"context"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllLogLogic {
	return GetAllLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllLogLogic) GetAllLog(req types.LogListReq) (interface{}, error) {
	data, err := l.svcCtx.LogModel.Find("", req.Start, req.End, req.Ip, req.Current, req.PageSize)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库查询失败", err)
	}
	total, err := l.svcCtx.LogModel.FindCount()
	return respx.NewListResponse(err, data, total, req.PageSize)
}
