package user

import (
	"context"
	"strings"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/common/tool"
	"workbench-server/common/utils"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutCurrentUserMobileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutCurrentUserMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) PutCurrentUserMobileLogic {
	return PutCurrentUserMobileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutCurrentUserMobileLogic) PutCurrentUserMobile(req *types.UserPutMobileReq) (any, error) {
	uid, err := tool.GetJwtUid(l.ctx)
	if err != nil {
		return nil, errorx.ExternalError("手机号密码更改失败！", err)
	}
	//校验手机号
	UserMobile := strings.TrimSpace(req.UserMobile)
	if err = utils.VerifyMobileFormat(UserMobile); err != nil {
		return nil, errorx.ExternalError(err.Error(), nil)
	}
	if err = l.svcCtx.UserModel.UpdateMobile(uid, UserMobile); err != nil {
		return respx.Msg(err, nil, "手机号更改成功", "手机号更改失败")
	}
	return respx.Msg(err, nil, "手机号更改成功", "手机号更改失败")
}
