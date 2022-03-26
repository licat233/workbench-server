package user

import (
	"context"
	"net/http"

	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/common/filex"
	"workbench-server/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutCurrentUserAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutCurrentUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) PutCurrentUserAvatarLogic {
	return PutCurrentUserAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutCurrentUserAvatarLogic) PutCurrentUserAvatar(r *http.Request) (any, error) {
	path, err := filex.SaveFile(r, l.svcCtx.Config.Public.Upload.Images.Avatar.Path, true)
	return respx.NewDefaultResponse(err, path)
}
