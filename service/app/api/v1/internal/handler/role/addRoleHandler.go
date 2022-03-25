package role

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"workbench-server/service/app/api/v1/internal/logic/role"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"
)

func AddRoleHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewAddRoleLogic(r.Context(), ctx)
		resp, err := l.AddRole(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
