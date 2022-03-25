package global

import (
	"errors"
	"net/http"
	"workbench-server/common/errorx"
	"workbench-server/common/utils"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
)

var (
	CasbinObj *casbin.Enforcer
)

func CasbinInitialize() {
	a, err := gormadapter.NewAdapter("mysql", "workbench:8kMRrkDB2aKkf5Mk@tcp(0.0.0.0:3306)/workbench?parseTime=True&charset=utf8mb4&collation=utf8mb4_unicode_ci", true)
	if err != nil {
		logx.Errorf("连接数据库错误: %v", err)
		return
	}
	e, err := casbin.NewEnforcer("/www/grpc/workbench/common/config/rbac_with_domains_model.conf", a)
	if err != nil {
		logx.Errorf("初始化casbin错误: %v", err)
		return
	}
	err = e.LoadPolicy()
	CasbinObj = e
}

func GetCasbinEnforcer() (*casbin.Enforcer, error) {
	a, err := gormadapter.NewAdapter("mysql", "workbench:8kMRrkDB2aKkf5Mk@tcp(0.0.0.0:3306)/workbench?parseTime=True&charset=utf8mb4&collation=utf8mb4_unicode_ci", true)
	if err != nil {
		logx.Errorf("连接数据库错误: %v", err)
		return nil, err
	}
	e, err := casbin.NewEnforcer("/www/grpc/workbench/common/config/rbac_with_domains_model.conf", a)
	if err != nil {
		logx.Errorf("初始化casbin错误: %v", err)
		return nil, err
	}
	err = e.LoadPolicy()
	return e, err
}

func Authorize(w http.ResponseWriter, r *http.Request) bool {
	// 获取用户角色ID
	sub, err := utils.GetJwtRoleId(r.Context())
	if err != nil {
		httpx.Error(w, err)
		return false
	}
	if sub == 1 { //超级管理员角色ID，直接通过
		return true
	}
	// 获取请求的URI
	obj := r.URL.RequestURI()
	// 获取请求方法
	act := r.Method
	// 获取用户ID
	uid, err := utils.GetJwtUid(r.Context())
	if err != nil {
		httpx.Error(w, err)
		return false
	}
	// 获取用户公司ID
	dom, err := utils.GetJwtCompanyId(r.Context())
	if err != nil {
		httpx.Error(w, err)
		return false
	}
	//判断策略中是否存在，sub角色-dom域(公司)-obj资源(url)-act操作(访问方法)
	if ok, _ := CasbinObj.Enforce(sub, dom, obj, act); !ok {
		logx.Errorf("很遗憾,权限验证没有通过，用户[%d]通过[%s]方法访问[%s]", uid, act, obj)
		resp := errorx.NewRequestError("权限不足", errors.New("当前用户无权限访问"))
		httpx.Error(w, resp)
		return false
	}
	return true
}
