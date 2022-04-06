package user

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"
	"workbench-server/service/v1/app/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) PutUserLogic {
	return PutUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutUserLogic) PutUser(req *types.UserInfoReq) (any, error) {
	// todo: 功能未验证
	data := &model.User{
		Uid:             req.Uid,
		UserRealName:    req.RealName,
		UserAvatar:      req.Avatar,
		UserGender:      req.Gender,
		Passport:        req.Passport,
		Password:        req.Password,
		UserEmail:       req.Email,
		UserMobile:      req.Mobile,
		UserRoleId:      req.RoleId,
		UserCompanyId:   req.CompanyId,
		UserQq:          req.QQ,
		UserWechat:      req.Wechat,
		UserStatus:      req.Status,
		UserDescription: req.Description,
	}
	err := l.svcCtx.UserModel.Update(data)
	return respx.Msg(err, nil, "更新成功", "更新失败")
}
