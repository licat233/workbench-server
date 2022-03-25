package rbac

import (
	"net/http"

	"workbench-server/service/app/api/v1/internal/logic/rbac"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAllActionsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := rbac.NewGetAllActionsLogic(r.Context(), ctx)
		resp, err := l.GetAllActions()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
