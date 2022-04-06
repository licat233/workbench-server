package article

import (
	"net/http"

	"workbench-server/service/v1/app/api/internal/logic/article"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ArticleListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := article.NewArticleListLogic(r.Context(), ctx)
		resp, err := l.ArticleList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
