package rbac

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"workbench-server/service/app/api/v1/internal/logic/rbac"
	"workbench-server/service/app/api/v1/internal/svc"
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
