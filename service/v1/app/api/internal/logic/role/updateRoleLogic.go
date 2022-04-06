package role

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"
	"workbench-server/service/v1/app/model"

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

func (l *UpdateRoleLogic) UpdateRole(req *types.RoleReq) (any, error) {
	data := &model.SysRole{
		Rid:             req.Rid,
		RoleName:        req.Name,
		RoleStatus:      req.Status,
		RoleDescription: req.Description,
		RoleIcon:        req.Icon,
	}
	err := l.svcCtx.RoleModel.Update(data)
	return respx.DefaultSingleResp(err, data)
}
