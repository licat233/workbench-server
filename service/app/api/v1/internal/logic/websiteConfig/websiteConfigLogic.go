package websiteConfig

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebsiteConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWebsiteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) WebsiteConfigLogic {
	return WebsiteConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebsiteConfigLogic) WebsiteConfig() (interface{}, error) {
	data, err := l.svcCtx.WebsiteConfigModel.FindOne(1)
	return respx.NewDefaultResponse(err, data)
}
