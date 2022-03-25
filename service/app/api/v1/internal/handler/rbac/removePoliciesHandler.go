package rbac

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"workbench-server/service/app/api/v1/internal/logic/rbac"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"
)

func RemovePoliciesHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PoliciesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := rbac.NewRemovePoliciesLogic(r.Context(), ctx)
		resp, err := l.RemovePolicies(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
