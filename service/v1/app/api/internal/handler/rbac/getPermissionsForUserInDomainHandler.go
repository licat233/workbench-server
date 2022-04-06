package rbac

import (
	"net/http"

	"workbench-server/service/v1/app/api/internal/logic/rbac"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPermissionsForUserInDomainHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleAndDomainReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := rbac.NewGetPermissionsForUserInDomainLogic(r.Context(), ctx)
		resp, err := l.GetPermissionsForUserInDomain(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
