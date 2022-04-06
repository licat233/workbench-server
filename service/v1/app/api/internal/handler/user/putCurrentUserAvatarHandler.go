package user

import (
	"net/http"

	"workbench-server/service/v1/app/api/internal/logic/user"
	"workbench-server/service/v1/app/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PutCurrentUserAvatarHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewPutCurrentUserAvatarLogic(r.Context(), ctx)
		resp, err := l.PutCurrentUserAvatar(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
