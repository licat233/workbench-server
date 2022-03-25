package menu

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"workbench-server/service/app/api/v1/internal/logic/menu"
	"workbench-server/service/app/api/v1/internal/svc"
)

func UpdateBaseMenuHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewUpdateBaseMenuLogic(r.Context(), ctx)
		resp, err := l.UpdateBaseMenu()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
