package middleware

import (
	"net/http"
	"workbench-server/common/global"
	"workbench-server/config"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var CorsHandler rest.Middleware = func(next http.HandlerFunc) http.HandlerFunc {
	// 开发模式，放行全部
	if global.CONFIG.CORS.Mode == "Dev" || global.CONFIG.CORS.Mode == "dev" {
		return CorsAll(next)
	}
	return CorsByConfig(next)
}

// CorsByConfig 跨域处理中间件,按照配置处理跨域请求
func CorsByConfig(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 所有OPTIONS方法返回204状态
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			httpx.Error(w, nil)
			return
		}

		respHeader := w.Header()
		whitelist := checkCors(r.Header.Get("origin"))

		//如果whitelist为空，那就无法通过验证，直接返回204状态
		if whitelist == nil {
			w.WriteHeader(http.StatusNonAuthoritativeInfo)
			httpx.Error(w, nil)
			return
		}

		// 通过检查, 添加请求头
		respHeader.Add("Access-Control-Allow-Origin", whitelist.AllowOrigin)
		respHeader.Add("Access-Control-Allow-Headers", whitelist.AllowHeaders)
		respHeader.Add("Access-Control-Allow-Methods", whitelist.AllowMethods)
		respHeader.Add("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
		if whitelist.AllowCredentials {
			respHeader.Add("Access-Control-Allow-Credentials", "true")
		}

		// 处理请求
		next(w, r)
	}
}

// CorsAll 直接放行所有跨域请求并放行所有 OPTIONS 方法
func CorsAll(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 所有OPTIONS方法返回204状态
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			httpx.Error(w, nil)
			return
		}
		respHeader := w.Header()
		origin := r.Header.Get("Origin") //允许请求源
		respHeader.Add("Access-Control-Allow-Origin", origin)
		respHeader.Add("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, X-Token, X-User-Id")
		respHeader.Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		respHeader.Add("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		respHeader.Add("Access-Control-Allow-Credentials", "true")

		// 处理请求
		next(w, r)
	}
}

// checkCors 从配置文件中加载当前源的配置
func checkCors(currentOrigin string) *config.CORSWhitelist {
	for _, whitelist := range global.CONFIG.CORS.Whitelist {
		// 遍历配置中的跨域头，寻找匹配项
		if currentOrigin == whitelist.AllowOrigin {
			return &whitelist
		}
	}
	return nil
}
