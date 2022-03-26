package rbac

import (
	"net/http"

	"workbench-server/app/api/v1/internal/logic/rbac"
	"workbench-server/app/api/v1/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPolicyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := rbac.NewGetPolicyLogic(r.Context(), ctx)
		resp, err := l.GetPolicy()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
