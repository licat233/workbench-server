package user

import (
	"net/http"

	"workbench-server/service/v1/app/api/internal/logic/user"
	"workbench-server/service/v1/app/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DelCurrentUserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewDelCurrentUserInfoLogic(r.Context(), ctx)
		resp, err := l.DelCurrentUserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
