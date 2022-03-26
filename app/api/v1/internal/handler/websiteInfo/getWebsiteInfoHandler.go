package websiteInfo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"workbench-server/app/api/v1/internal/logic/websiteInfo"
	"workbench-server/app/api/v1/internal/svc"
)

func GetWebsiteInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := websiteInfo.NewGetWebsiteInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetWebsiteInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
