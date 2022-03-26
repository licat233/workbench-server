package captcha

import (
	"context"
	"image/color"

	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"
	"workbench-server/common/respx"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
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

func (l *MathCaptchaLogic) MathCaptcha(req *types.CaptchaReq) (interface{}, error) {
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
	var store = base64Captcha.DefaultMemStore
	driver := param.DriverMath.ConvertFonts()
	//目前仅支持math验证，后期如果需要，再开启多个方式
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

	c := base64Captcha.NewCaptcha(driver, store)
	key, base64Image, err := c.Generate()
	return respx.NewDefaultResponse(err, &types.CaptchaInfo{
		CaptchaKey:  key,
		Base64Image: base64Image,
	})
}
