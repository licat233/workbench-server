package user

import (
	"context"
	"net/http"

	"workbench-server/common/filex"
	"workbench-server/common/global"
	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"

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
	path, err := filex.SaveHttpImg(r, global.CONFIG.StaticConfig.Root.Upload.Images.Avatar.Path, "avatar")
	return respx.DefaultSingleResp(err, path)
}
