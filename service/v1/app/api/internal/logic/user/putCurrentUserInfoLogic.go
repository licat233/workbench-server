package user

import (
	"context"
	"errors"

	"workbench-server/common/respx"
	"workbench-server/common/tool"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"
	"workbench-server/service/v1/app/model"

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

func (l *PutCurrentUserInfoLogic) PutCurrentUserInfo(req *types.UserSelfInfoReq) (any, error) {
	uid, err := tool.GetJwtUid(l.ctx)
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
	return respx.Msg(err, nil, "更新用户信息成功！", "更新用户信息失败！")
}
