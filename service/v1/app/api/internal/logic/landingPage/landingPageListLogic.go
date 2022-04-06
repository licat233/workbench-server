package landingPage

import (
	"context"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LandingPageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLandingPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) LandingPageListLogic {
	return LandingPageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LandingPageListLogic) LandingPageList(req *types.ListReq) (any, error) {
	data, err := l.svcCtx.LandingPageModel.FindAll(req.Current, req.PageSize)
	if err != nil {
		return nil, errorx.InternalError("数据库查询失败", err)
	}
	total, err := l.svcCtx.LandingPageModel.FindCount()
	return respx.DefaultListResp(err, data, total, req.PageSize, req.Current)
}
