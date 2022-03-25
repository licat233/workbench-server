package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"workbench-server/common/utils"
	"workbench-server/service/app/api/v1/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/dgrijalva/jwt-go"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type JwtMiddleware struct {
	Config     config.Config
	SigningKey []byte
	RedisConn  *redis.Redis
}

func NewJwtMiddleware(c config.Config) *JwtMiddleware {
	return &JwtMiddleware{
		Config:     c,
		SigningKey: []byte(c.Auth.AccessSecret),
		RedisConn:  redis.New(c.Redis[0].Host, redis.WithPass(c.Redis[0].Pass)),
	}
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get("Authorization")
		TokenKey := fmt.Sprintf("token:blacklist:%s", jwtToken)
		res, err := m.RedisConn.Get(TokenKey)
		if err != nil {
			//不存在该token
			httpx.Error(w, errors.New("身份验证失败:未知身份"))
			return
		}
		if res == "logout" {
			httpx.Error(w, errors.New("当前登录已被注销"))
			return
		}
		platform := fmt.Sprintf("%v", r.Context().Value("platform"))
		if res == "repeal" {
			httpx.Error(w, fmt.Errorf("当前账号已在其它%s设备上登录，被迫下线", platform))
			return
		}
		//设备验证
		currentPlatform, err := utils.RemotePlatform(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if currentPlatform != platform {
			httpx.Error(w, errors.New("身份验证失败:设备验证失败"))
			return
		}
		jwtIp := r.Context().Value("ip")
		ip := fmt.Sprintf("%v", jwtIp)
		//ip 校验
		if utils.RemoteIp(r) != ip {
			httpx.Error(w, errors.New("身份验证失败"))
			return
		}
		// Passthrough to next handler if need
		next(w, r)
	}
}

// 生成token
func (m *JwtMiddleware) GenerateJwtToken(claims jwt.MapClaims) (tokenString string, err error) {
	// 创建一个新的令牌对象，指定签名方法和声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密码签名并获得完整的编码令牌作为字符串
	return token.SignedString(m.SigningKey)
}

// 解析token
func (m *JwtMiddleware) ParseToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return m.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

//GetIdFromClaims 从token对象里获得参数(key)对应的值
// 示例 ：GetIdFromClaims("username", token.claims) 其中token是已经解密的token
func (m *JwtMiddleware) GetIdFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	//确保是map类型
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)
			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}
