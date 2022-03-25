package project

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"workbench-server/service/app/api/v1/internal/logic/project"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"
)

func GetProjectListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := project.NewGetProjectListLogic(r.Context(), ctx)
		resp, err := l.GetProjectList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
