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

func (l *PutCurrentUserEmailLogic) PutCurrentUserEmail(req *types.UserPutEmailReq) (any, error) {
	uid, err := tool.GetJwtUid(l.ctx)
	if err != nil {
		return nil, errorx.ExternalError("邮箱密码更改失败！", err)
	}
	//校验邮箱
	UserEmail := strings.TrimSpace(req.UserEmail)
	if err = utils.VerifyEmailFormat(UserEmail); err != nil {
		return nil, errorx.ExternalError(err.Error(), nil)
	}
	err = l.svcCtx.UserModel.UpdateEmail(uid, UserEmail)
	return respx.Msg(err, nil, "邮箱更改成功", "邮箱更改失败")
}
