package common

import (
	"context"
	"net/http"

	"workbench-server/common/filex"
	"workbench-server/common/respx"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) UploadImgLogic {
	return UploadImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImgLogic) UploadImg(r *http.Request) (any, error) {
	path, err := filex.SaveFile(r, l.svcCtx.Config.Storage.Upload.Images.Avatar.Path, true)
	return respx.NewDefaultResponse(err, path)
}
