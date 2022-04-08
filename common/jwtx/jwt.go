package jwtx

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"workbench-server/common/global"
	"workbench-server/common/utils"

	"github.com/dgrijalva/jwt-go"
	// uuid "github.com/satori/go.uuid"
)

var DefaultDuration int64 = 86400 //一天时间

// type JWT struct {
// 	Blacklist model.JwtBlacklistsModel
// 	Whitelist model.JwtWhitelistsModel
// }

type BaseClaims struct {
	// UUID      uuid.UUID
	Uid      int64  //用户ID
	Ip       string //设备IP
	Platform string //设备平台
	Duration int64  //有效时间
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func CreateMapClaims(baseClaims BaseClaims) jwt.MapClaims {
	claims := make(jwt.MapClaims)
	now := time.Now().Unix()
	claims["exp"] = now + baseClaims.Duration //过期时间
	claims["iat"] = now                       //生效时间
	claims["nbf"] = now
	claims["uid"] = baseClaims.Uid           //用户id
	claims["ip"] = baseClaims.Ip             //用户ip
	claims["platform"] = baseClaims.Platform //用户设备
	return claims
}

func CreateToken(claims jwt.MapClaims) (string, error) {
	// 创建一个新的令牌对象，指定签名方法和声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密码签名并获得完整的编码令牌作为字符串
	return token.SignedString([]byte(global.CONFIG.JWT.AccessSecret))
}

// 解析token
func ParseToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(global.CONFIG.JWT.AccessSecret), nil
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

//GetClaims 从请求中获取claims
func GetClaims(r *http.Request) (*jwt.MapClaims, error) {
	token := r.Header.Get("x-token")
	claims, err := ParseToken(token)
	if err != nil {
		global.ZAP.Error("从request中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

//GetClaimsValue 从token对象里获得参数(key)对应的值
// 示例 GetClaimsValue("username", token.claims) 其中token是已经解密的token
func GetClaimsValue(key string, claims jwt.MapClaims) string {
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

func ExpiredTime(claims jwt.MapClaims) int64 {
	s := GetClaimsValue("exp", claims)
	if len(s) == 0 {
		return 0
	}
	res, err := utils.ExpToNumber(s)
	if err != nil {
		return 0
	}
	return res
}
