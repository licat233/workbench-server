package teacher

import (
	"context"

	"workbench-server/global/errorx"
	"workbench-server/global/respx"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTeacherListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetTeacherListLogic {
	return GetTeacherListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTeacherListLogic) GetTeacherList(req types.ListReq) (interface{}, error) {
	data, err := l.svcCtx.TeacherModel.FindAll(req.Current, req.PageSize)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库查询失败", err)
	}
	total, err := l.svcCtx.TeacherModel.FindCount()
	return respx.NewListResponse(err, data, total, req.PageSize)
}
