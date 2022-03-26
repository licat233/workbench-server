package article

import (
	"net/http"

	"workbench-server/app/api/v1/internal/logic/article"
	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"

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
