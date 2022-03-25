package log

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"workbench-server/service/app/api/v1/internal/logic/log"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"
)

func GetAllLogHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := log.NewGetAllLogLogic(r.Context(), ctx)
		resp, err := l.GetAllLog(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
