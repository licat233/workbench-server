package captcha

import (
	"context"

	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

type ProcessCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcessCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProcessCaptchaLogic {
	return ProcessCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcessCaptchaLogic) ProcessCaptcha(req *types.ProcessCaptchaReq) (any, error) {
	ok := base64Captcha.DefaultMemStore.Verify(req.CaptchaKey, req.Solution, true)
	if ok {
		return respx.SingleResp("验证成功", ok), nil
	}
	return respx.SingleResp("验证失败", ok), nil
}
