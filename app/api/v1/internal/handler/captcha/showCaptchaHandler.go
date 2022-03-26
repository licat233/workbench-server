package captcha

import (
	"net/http"

	"workbench-server/app/api/v1/internal/logic/captcha"
	"workbench-server/app/api/v1/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShowCaptchaHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := captcha.NewShowCaptchaLogic(r.Context(), ctx)
		err := l.ShowCaptcha(w)
		if err != nil {
			httpx.Error(w, err)
		}
		// else {
		//以为logic内已经调用过一次 response.WriteHeader 方法，再次调次调用，会触发打印：
		//http: superfluous response.WriteHeader
		// 	httpx.Ok(w)
		// }
	}
}
