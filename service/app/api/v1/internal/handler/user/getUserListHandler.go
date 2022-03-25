package user

import (
	"net/http"

	"workbench-server/service/app/api/v1/internal/logic/user"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetUserListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewGetUserListLogic(r.Context(), ctx)
		resp, err := l.GetUserList(r, req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
