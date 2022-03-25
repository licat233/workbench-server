package common

import (
	"net/http"

	"workbench-server/service/app/api/v1/internal/logic/common"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadImgHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := common.NewUploadImgLogic(r.Context(), ctx)
		resp, err := l.UploadImg(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
