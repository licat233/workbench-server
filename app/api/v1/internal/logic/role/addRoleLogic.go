package role

import (
	"context"

	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"
	"workbench-server/app/model"
	"workbench-server/common/respx"

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

func (l *AddRoleLogic) AddRole(req *types.RoleReq) (interface{}, error) {
	data := &model.SysRole{
		Rid:             req.Rid,
		RoleName:        req.Name,
		RoleStatus:      req.Status,
		RoleDescription: req.Description,
		RoleIcon:        req.Icon,
	}
	_, err := l.svcCtx.RoleModel.Insert(data)
	return respx.NewDefaultResponse(err, nil, "添加成功", "添加失败")
}
