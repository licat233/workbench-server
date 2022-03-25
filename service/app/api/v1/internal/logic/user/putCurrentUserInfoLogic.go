package user

import (
	"context"
	"errors"

	"workbench-server/common/respx"
	"workbench-server/common/utils"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"
	"workbench-server/service/app/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutCurrentUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutCurrentUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) PutCurrentUserInfoLogic {
	return PutCurrentUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutCurrentUserInfoLogic) PutCurrentUserInfo(req types.UserSelfInfoReq) (interface{}, error) {
	uid, err := utils.GetJwtUid(l.ctx)
	if err != nil {
		return nil, errors.New("更新用户信息失败！因为获取不到你的uid")
	}
	var user = &model.UserPutData{
		Uid:             uid,
		UserRealName:    req.RealName,
		UserRoleId:      req.RoleId,
		UserGender:      req.Gender,
		UserQq:          req.QQ,
		UserWechat:      req.Wechat,
		UserDescription: req.Description,
	}
	// utils.StructAssign(user, &req)
	// user.Uid = uid
	err = l.svcCtx.UserModel.BaseUpdate(user)
	return respx.NewDefaultResponse(err, nil, "更新用户信息成功！", "更新用户信息失败！")
}
