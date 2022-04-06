package log

import (
	"context"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

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

func (l *GetAllLogLogic) GetAllLog(req *types.LogListReq) (any, error) {
	data, err := l.svcCtx.LogModel.Find("", req.Start, req.End, req.Ip, req.Current, req.PageSize)
	if err != nil {
		return nil, errorx.InternalError("数据库查询失败", err)
	}
	total, err := l.svcCtx.LogModel.FindCount()
	return respx.DefaultListResp(err, data, total, req.PageSize, req.Current)
}
