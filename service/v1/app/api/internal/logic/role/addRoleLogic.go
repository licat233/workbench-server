package role

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"
	"workbench-server/service/v1/app/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddRoleLogic {
	return AddRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRoleLogic) AddRole(req *types.RoleReq) (any, error) {
	data := &model.SysRole{
		Rid:             req.Rid,
		RoleName:        req.Name,
		RoleStatus:      req.Status,
		RoleDescription: req.Description,
		RoleIcon:        req.Icon,
	}
	_, err := l.svcCtx.RoleModel.Insert(data)
	return respx.DefaultStateResp(err)
}
