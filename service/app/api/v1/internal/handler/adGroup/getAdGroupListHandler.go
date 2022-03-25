package adGroup

import (
	"net/http"

	"workbench-server/service/app/api/v1/internal/logic/adGroup"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAdGroupListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := adGroup.NewGetAdGroupListLogic(r.Context(), ctx)
		resp, err := l.GetAdGroupList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
