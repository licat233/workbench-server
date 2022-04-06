package websiteInfo

import (
	"net/http"

	"workbench-server/service/v1/app/api/internal/logic/websiteInfo"
	"workbench-server/service/v1/app/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
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
