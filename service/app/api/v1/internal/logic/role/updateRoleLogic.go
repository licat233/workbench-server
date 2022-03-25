package role

import (
	"context"

	"workbench-server/global/respx"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"
	"workbench-server/service/app/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateRoleLogic {
	return UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req types.RoleReq) (interface{}, error) {
	data := &model.SysRole{
		Rid:             req.Rid,
		RoleName:        req.Name,
		RoleStatus:      req.Status,
		RoleDescription: req.Description,
		RoleIcon:        req.Icon,
	}
	err := l.svcCtx.RoleModel.Update(data)
	return respx.NewDefaultResponse(err, data, "更新数据成功", "更新数据失败")
}
