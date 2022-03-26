package websiteInfo

import (
	"context"

	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/common/respx"

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

func (l *GetWebsiteInfoLogic) GetWebsiteInfo() (interface{}, error) {
	data, err := l.svcCtx.WebsiteConfigModel.FindOne(1)
	return respx.NewDefaultResponse(err, data)
}
