package ad

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"workbench-server/app/api/v1/internal/logic/ad"
	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"
)

func GetAdCampaignListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := ad.NewGetAdCampaignListLogic(r.Context(), svcCtx)
		resp, err := l.GetAdCampaignList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
