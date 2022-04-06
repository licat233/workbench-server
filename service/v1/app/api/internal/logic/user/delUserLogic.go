package user

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *DelUserLogic) DelUser(req *types.BatchIdReq) (any, error) {
	var err error
	for _, uid := range req.Id {
		err = l.svcCtx.UserModel.Delete(uid)
		if err != nil {
			break
		}
	}
	return respx.Msg(err, nil, "删除成功", "删除失败")
}
