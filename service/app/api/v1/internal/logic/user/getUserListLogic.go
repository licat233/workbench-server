package user

import (
	"context"
	"net/http"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserListLogic {
	return GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(r *http.Request, req types.ListReq) (interface{}, error) {
	total, err := l.svcCtx.UserModel.FindCount()
	if err != nil {
		return nil, errorx.NewDefaultError("数据库查询失败", err)
	}
	data, err := l.svcCtx.UserModel.FindAll(r, req.Current, req.PageSize)
	return respx.NewListResponse(err, data, total, req.PageSize)
}
