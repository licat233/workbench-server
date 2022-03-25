package user

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"workbench-server/service/app/api/v1/internal/logic/user"
	"workbench-server/service/app/api/v1/internal/svc"
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
