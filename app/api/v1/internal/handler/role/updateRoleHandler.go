package role

import (
	"net/http"

	"workbench-server/app/api/v1/internal/logic/role"
	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateRoleHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewUpdateRoleLogic(r.Context(), ctx)
		resp, err := l.UpdateRole(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
