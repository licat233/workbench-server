package rbac

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"workbench-server/service/app/api/v1/internal/logic/rbac"
	"workbench-server/service/app/api/v1/internal/svc"
)

func GetAllObjectsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := rbac.NewGetAllObjectsLogic(r.Context(), ctx)
		resp, err := l.GetAllObjects()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
