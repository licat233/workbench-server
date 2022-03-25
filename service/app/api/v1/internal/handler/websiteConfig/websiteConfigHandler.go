package websiteConfig

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"workbench-server/service/app/api/v1/internal/logic/websiteConfig"
	"workbench-server/service/app/api/v1/internal/svc"
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
