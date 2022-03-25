package menu

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"workbench-server/service/app/api/v1/internal/logic/menu"
	"workbench-server/service/app/api/v1/internal/svc"
)

func GetBaseMenuByIdHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetBaseMenuByIdLogic(r.Context(), ctx)
		resp, err := l.GetBaseMenuById()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
