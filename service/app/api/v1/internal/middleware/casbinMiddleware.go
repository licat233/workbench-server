package middleware

import (
	"errors"
	"net/http"
	"workbench-server/common/errorx"
	"workbench-server/common/global"
	"workbench-server/common/utils"

	"github.com/casbin/casbin/v2"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
)

type CasbinMiddleware struct {
	CasbinObj *casbin.Enforcer
}

func NewCasbinMiddleware() *CasbinMiddleware {
	return &CasbinMiddleware{
		CasbinObj: global.CasbinObj,
	}
}

func (m *CasbinMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if true { //暂时全部开放，在RBAC权限系统未完善之前
			next(w, r)
			return
		}

		// 获取用户角色ID
		sub, err := utils.GetJwtRoleId(r.Context())
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if sub == 1 { //超级管理员角色ID，直接通过
			next(w, r)
			return
		}

		// 获取请求的URI
		obj := r.URL.RequestURI()
		// 获取请求方法
		act := r.Method
		// 获取用户ID
		uid, err := utils.GetJwtUid(r.Context())
		if err != nil {
			httpx.Error(w, err)
			return
		}
		// 获取用户公司ID
		dom, err := utils.GetJwtCompanyId(r.Context())
		if err != nil {
			httpx.Error(w, err)
			return
		}
		//判断策略中是否存在，sub角色-dom域(公司)-obj资源(url)-act操作(访问方法)
		if ok, _ := m.CasbinObj.Enforce(sub, dom, obj, act); !ok {
			logx.Errorf("很遗憾,权限验证没有通过，用户[%d]通过[%s]方法访问[%s]", uid, act, obj)
			resp := errorx.NewRequestError("权限不足", errors.New("当前用户无权限访问"))
			httpx.Error(w, resp)
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
