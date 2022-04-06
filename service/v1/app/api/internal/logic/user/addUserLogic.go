package user

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"
	"workbench-server/service/v1/app/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddUserLogic {
	return AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.UserAddReq) (any, error) {
	data := &model.User{
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
	_, err := l.svcCtx.UserModel.Insert(data)
	return respx.Msg(err, nil, "添加成功", "添加失败")
}
