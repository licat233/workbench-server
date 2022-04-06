package storage

import (
	"io/ioutil"
	"net/http"

	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		body, err := ioutil.ReadFile(req.File)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		w.Write(body)
		// l := storage.NewDownloadLogic(r.Context(), ctx)
		// err := l.Download(&req)
		// if err != nil {
		// 	httpx.Error(w, err)
		// } else {
		// 	httpx.Ok(w)
		// }
	}
}
