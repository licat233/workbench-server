package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"workbench-server/common/errorx"
	"workbench-server/common/global"
	"workbench-server/common/tool"
	"workbench-server/service/v1/app/model"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type CasbinMiddleware struct {
	CasbinService *casbin.SyncedEnforcer
	UserModel     model.UserModel
}

func NewCasbinMiddleware() *CasbinMiddleware {
	return &CasbinMiddleware{
		CasbinService: global.CASBIN,
		UserModel:     model.NewUserModel(global.SQLConn, global.CONFIG.Cache.Redis),
	}
}

func (m *CasbinMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if true { //暂时全部开放，在RBAC权限系统未完善之前
			next(w, r)
			return
		}

		// 获取用户ID
		uid, err := tool.GetJwtUid(r.Context())
		if err != nil {
			httpx.Error(w, err)
			return
		}
		//获取用户信息
		userInfo, err := m.UserModel.FindOne(uid)
		if err == model.ErrNotFound {
			resp := errorx.AuthError("权限不足", errors.New("当前用户不存在数据库中"))
			httpx.Error(w, resp)
		}
		if err != nil {
			resp := errorx.InternalError("系统异常，暂时无法访问，请联系管理员", errors.New("当前用户不存在数据库中"))
			httpx.Error(w, resp)
		}

		// 获取用户角色ID
		sub := userInfo.UserRoleId
		if sub == 1 { //超级管理员角色ID，直接通过
			next(w, r)
			return
		}

		// 获取请求的URI
		obj := r.URL.RequestURI()
		// 获取请求方法
		act := r.Method

		// 获取用户公司ID
		dom := userInfo.UserCompanyId

		//校验casbin是否已初始化，防止报Panic
		if m.CasbinService == nil {
			resp := errorx.ServerError("系统异常，暂时无法访问，请联系管理员", errors.New("服务器RBAC系统异常"))
			httpx.Error(w, resp)
			return
		}

		//判断策略中是否存在，sub角色(职位)-dom域(公司)-obj资源(url)-act操作(访问方法)
		if ok, _ := m.CasbinService.Enforce(sub, dom, obj, act); !ok {
			global.ZAP.Error(fmt.Sprintf("很遗憾,权限验证没有通过，用户[%d]通过[%s]方法访问[%s]", uid, act, obj))
			resp := errorx.AuthError("权限不足", errors.New("当前用户无权限访问"))
			httpx.Error(w, resp)
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
