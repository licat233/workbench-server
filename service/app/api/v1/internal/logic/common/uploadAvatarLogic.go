package common

import (
	"context"
	"net/http"

	"workbench-server/common/filex"
	"workbench-server/common/respx"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UploadAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) UploadAvatarLogic {
	return UploadAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadAvatarLogic) UploadAvatar(r *http.Request) (any, error) {
	path, err := filex.SaveFile(r, l.svcCtx.Config.Storage.Upload.Images.Avatar.Path, true)
	return respx.NewDefaultResponse(err, path)
}
