package captcha

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"

	"workbench-server/common/errorx"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

type ShowCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) ShowCaptchaLogic {
	return ShowCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowCaptchaLogic) ShowCaptcha(w http.ResponseWriter) error {
	param := &CaptchaConfig{
		CaptchaType: "math",
		DriverMath: &base64Captcha.DriverMath{
			Fonts:           []string{"actionj.ttf", "3Dumb.ttf"},
			Height:          60,
			NoiseCount:      0,
			ShowLineOptions: 6,
			Width:           200,
		},
		Id: "",
	}

	var store = base64Captcha.DefaultMemStore
	driver := param.DriverMath.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	captchaId, b64s, err := c.Generate()
	if err != nil {
		return errorx.NewDefaultError("获取验证码失败", err)
	}
	mark := "data:image/png;base64,"
	m := strings.Index(b64s, mark)
	if m == -1 {
		return errorx.NewDefaultError("获取验证码失败", err)
	}
	n := m + len(mark)
	base := string([]byte(b64s)[n:])
	p, err := base64.StdEncoding.DecodeString(base)
	if err != nil {
		return errorx.NewDefaultError("获取验证码失败", err)
	}
	w.Header().Set("cache-control", "no-cache")
	w.Header().Set("content-type", "image/png")
	w.Header().Add("captcha-key", captchaId)
	_, err = w.Write(p)
	if err != nil {
		logx.Errorf("http写入验证码图片数据失败", err)
		// return errorx.NewDefaultError("获取验证码失败", err)
	}
	// lastLog, err := l.svcCtx.LogModel.FindLastLogin(2, "mac")
	// fmt.Println(lastLog, err)
	return nil
}
