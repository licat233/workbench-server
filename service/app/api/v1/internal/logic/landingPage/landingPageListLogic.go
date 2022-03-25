package landingPage

import (
	"context"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *LandingPageListLogic) LandingPageList(req types.ListReq) (interface{}, error) {
	data, err := l.svcCtx.LandingPageModel.FindAll(req.Current, req.PageSize)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库查询失败", err)
	}
	total, err := l.svcCtx.LandingPageModel.FindCount()
	return respx.NewListResponse(err, data, total, req.PageSize)
}
