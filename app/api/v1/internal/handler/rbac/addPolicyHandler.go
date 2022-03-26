package rbac

import (
	"net/http"

	"workbench-server/app/api/v1/internal/logic/rbac"
	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddPolicyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PolicyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := rbac.NewAddPolicyLogic(r.Context(), ctx)
		resp, err := l.AddPolicy(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
