package menu

import (
	"net/http"

	"workbench-server/service/v1/app/api/internal/logic/menu"
	"workbench-server/service/v1/app/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddBaseMenuHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewAddBaseMenuLogic(r.Context(), ctx)
		resp, err := l.AddBaseMenu()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
