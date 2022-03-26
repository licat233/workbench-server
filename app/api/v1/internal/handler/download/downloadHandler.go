package download

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"workbench-server/app/api/v1/internal/logic/download"
	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := download.NewDownloadLogic(r.Context(), svcCtx)
		err := l.Download(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
