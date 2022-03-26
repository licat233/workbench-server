package menu

import (
	"net/http"

	"workbench-server/app/api/v1/internal/logic/menu"
	"workbench-server/app/api/v1/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMenuAuthorityHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetMenuAuthorityLogic(r.Context(), ctx)
		resp, err := l.GetMenuAuthority()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
