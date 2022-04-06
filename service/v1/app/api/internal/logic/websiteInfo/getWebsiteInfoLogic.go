package websiteInfo

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebsiteInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWebsiteInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebsiteInfoLogic {
	return &GetWebsiteInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWebsiteInfoLogic) GetWebsiteInfo() (any, error) {
	data, err := l.svcCtx.WebsiteConfigModel.FindOne(1)
	return respx.DefaultSingleResp(err, data)
}
