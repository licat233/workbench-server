package user

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DelUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) DelUserLogic {
	return DelUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelUserLogic) DelUser(req types.BatchIdReq) (interface{}, error) {
	var err error
	for _, uid := range req.Id {
		err = l.svcCtx.UserModel.Delete(uid)
		if err != nil {
			break
		}
	}
	return respx.NewDefaultResponse(err, nil, "删除成功", "删除失败")
}
