package captcha

import (
	"context"
	"image/color"

	"workbench-server/common/captchax"
	"workbench-server/common/global"
	"workbench-server/common/respx"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

type MathCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMathCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) MathCaptchaLogic {
	return MathCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MathCaptchaLogic) MathCaptcha(req *types.CaptchaReq) (any, error) {
	param := &CaptchaConfig{
		CaptchaType: "math",
		DriverMath: &base64Captcha.DriverMath{
			BgColor:         (*color.RGBA)(&req.BgColor),
			Fonts:           []string{"actionj.ttf", "3Dumb.ttf"},
			Height:          req.Height,
			NoiseCount:      0,
			ShowLineOptions: 6,
			Width:           req.Width,
		},
	}
	driver := param.DriverMath.ConvertFonts()
	//目前仅支持math验证，后期如果有需要，再开启多个方式
	//switch in.CaptchaType {
	//case "audio":
	//	driver = param.DriverAudio
	//case "string":
	//	driver = param.DriverString.ConvertFonts()
	//case "math":
	//	driver = param.DriverMath.ConvertFonts()
	//case "chinese":
	//	driver = param.DriverChinese.ConvertFonts()
	//default:
	//	driver = param.DriverDigit
	//}

	c := base64Captcha.NewCaptcha(driver, captchax.Store)
	key, base64Image, err := c.Generate()
	if err != nil {
		global.ZAP.Error("验证码获取失败!", zap.Error(err))
	}
	return respx.DefaultSingleResp(err, &types.CaptchaInfo{
		CaptchaKey:  key,
		Base64Image: base64Image,
	})
}
