package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"workbench-server/common/errorx"
	"workbench-server/common/jwtx"
	"workbench-server/common/jwtx/model"
	"workbench-server/common/utils"

	"github.com/zeromicro/go-zero/rest/httpx"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type JwtMiddleware struct {
	JwtService *jwtx.JWT
}

func NewJwtMiddleware() *JwtMiddleware {
	return &JwtMiddleware{
		JwtService: jwtx.NewJWT(),
	}
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//检查是否存在于黑名单里面
		jwtToken := r.Header.Get("x-token")
		info, err := model.FindJwtInRedis(jwtToken)
		if err != nil {
			httpx.Error(w, errorx.InternalError("服务器错误", err))
			return
		}

		platform := fmt.Sprintf("%v", r.Context().Value("platform"))
		//jwt存在黑名单中
		if info != nil {
			switch info.Status {
			case model.LOGOUT:
				httpx.Error(w, errors.New("当前登录已被注销"))
				return
			case model.REPEAL:

				httpx.Error(w, fmt.Errorf("当前账号已在其它%s设备上登录，被迫下线", platform))
				return
			case model.OVERDUE:
				httpx.Error(w, errors.New("当前登录已过期"))
				return
			default:
				httpx.Error(w, errors.New("当前登录状态异常，请重新登录"))
				return
			}
		}

		//设备验证
		currentPlatform, err := utils.RemotePlatform(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if currentPlatform != platform {
			httpx.Error(w, errors.New("身份验证失败，身份异常"))
			return
		}

		currentIP := utils.RemoteIp(r)
		ip := fmt.Sprint(r.Context().Value("ip"))
		//ip 校验
		if currentIP != ip {
			httpx.Error(w, errors.New("身份验证失败，身份异常"))
			return
		}
		// Passthrough to next handler if need
		next(w, r)
	}
}
