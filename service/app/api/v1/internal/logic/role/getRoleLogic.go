package role

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRoleLogic {
	return GetRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleLogic) GetRole(req types.IdReq) (interface{}, error) {
	data, err := l.svcCtx.RoleModel.FindOne(req.Id)
	return respx.NewDefaultResponse(err, data, "获取数据成功", "获取数据失败")
}
