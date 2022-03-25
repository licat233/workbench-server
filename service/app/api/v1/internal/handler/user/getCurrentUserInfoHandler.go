package user

import (
	"net/http"

	"workbench-server/service/app/api/v1/internal/logic/user"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCurrentUserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetCurrentUserInfoLogic(r.Context(), ctx)
		resp, err := l.GetCurrentUserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
