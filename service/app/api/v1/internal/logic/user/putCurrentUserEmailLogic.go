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

type PutCurrentUserEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutCurrentUserEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) PutCurrentUserEmailLogic {
	return PutCurrentUserEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutCurrentUserEmailLogic) PutCurrentUserEmail(req types.UserPutEmailReq) (interface{}, error) {
	uid, err := utils.GetJwtUid(l.ctx)
	if err != nil {
		return nil, errorx.NewRequestError("邮箱密码更改失败！", err)
	}
	//校验邮箱
	UserEmail := strings.TrimSpace(req.UserEmail)
	if err = utils.VerifyEmailFormat(UserEmail); err != nil {
		return nil, errorx.NewRequestError(err.Error(), nil)
	}
	if err = l.svcCtx.UserModel.UpdateEmail(uid, UserEmail); err != nil {
		return respx.NewDefaultResponse(err, nil, "邮箱更改成功", "邮箱更改失败")
	}
	return respx.NewDefaultResponse(err, nil, "邮箱更改成功", "邮箱更改失败")
}
