package menu

import (
	"net/http"

	"workbench-server/service/app/api/v1/internal/logic/menu"
	"workbench-server/service/app/api/v1/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddMenuAuthorityHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewAddMenuAuthorityLogic(r.Context(), ctx)
		resp, err := l.AddMenuAuthority()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
