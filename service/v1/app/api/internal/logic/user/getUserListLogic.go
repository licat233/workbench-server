package user

import (
	"context"
	"net/http"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *GetUserListLogic) GetUserList(r *http.Request, req types.ListReq) (any, error) {
	total, err := l.svcCtx.UserModel.FindCount()
	if err != nil {
		return nil, errorx.InternalError("数据库查询失败", err)
	}
	data, err := l.svcCtx.UserModel.FindAll(r, req.Current, req.PageSize)
	return respx.DefaultListResp(err, data, total, req.PageSize, req.Current)
}
