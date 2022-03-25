package common

import (
	"net/http"

	"workbench-server/service/app/api/v1/internal/logic/common"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadAvatarHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := common.NewUploadAvatarLogic(r.Context(), ctx)
		resp, err := l.UploadAvatar(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
