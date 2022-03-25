package user

import (
	"context"
	"strings"

	"workbench-server/global/errorx"
	"workbench-server/global/respx"
	"workbench-server/global/utils"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

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

func (l *PutCurrentUserMobileLogic) PutCurrentUserMobile(req types.UserPutMobileReq) (interface{}, error) {
	uid, err := utils.GetJwtUid(l.ctx)
	if err != nil {
		return nil, errorx.NewRequestError("手机号密码更改失败！", err)
	}
	//校验手机号
	UserMobile := strings.TrimSpace(req.UserMobile)
	if err = utils.VerifyMobileFormat(UserMobile); err != nil {
		return nil, errorx.NewRequestError(err.Error(), nil)
	}
	if err = l.svcCtx.UserModel.UpdateMobile(uid, UserMobile); err != nil {
		return respx.NewDefaultResponse(err, nil, "手机号更改成功", "手机号更改失败")
	}
	return respx.NewDefaultResponse(err, nil, "手机号更改成功", "手机号更改失败")
}
