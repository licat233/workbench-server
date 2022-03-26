package websiteConfig

import (
	"net/http"

	"workbench-server/app/api/v1/internal/logic/websiteConfig"
	"workbench-server/app/api/v1/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func WebsiteConfigHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := websiteConfig.NewWebsiteConfigLogic(r.Context(), ctx)
		resp, err := l.WebsiteConfig()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
