package model

import (
	"encoding/json"
	"fmt"
	"workbench-server/common/errorx"
	"workbench-server/common/global"
)

var (
	REPEAL  = "repeal"  //设备单点登录机制，被迫失效
	LOGOUT  = "logout"  //被主动注销
	OVERDUE = "overdue" //已过期
)

type JwtBaseInfo struct {
	Uid    int64  `json:"uid"`    //用户ID
	Status string `json:"status"` //当前状态：REPEAL | LOGOUT | OVERDUE
}

//ExistInRedis 判断jwt是否存在黑名单中
func ExistInRedis(jwtToken string) (bool, error) {
	workbenchJwtBlacklistsJwtKey := fmt.Sprintf("%s%s", cacheWorkbenchJwtBlacklistsJwtPrefix, jwtToken)
	return global.RedisX.Exists(workbenchJwtBlacklistsJwtKey)
}

//FindJwtInRedis 在redis中查找jwt
func FindJwtInRedis(jwtToken string) (*JwtBaseInfo, error) {
	workbenchJwtBlacklistsJwtKey := fmt.Sprintf("%s%s", cacheWorkbenchJwtBlacklistsJwtPrefix, jwtToken)
	value, err := global.RedisX.Get(workbenchJwtBlacklistsJwtKey)
	if err != nil {
		return nil, err
	}
	if len(value) == 0 {
		return nil, nil
	}
	res := new(JwtBaseInfo)
	err = json.Unmarshal([]byte(value), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

/** InsertToRedis 插入jwt基本kv信息：
 * key: token
 * value: {uid int64, status string}
 */
func InsertToRedis(uid int64, jwt string, exp int64, status string) error {
	workbenchJwtBlacklistsJwtKey := fmt.Sprintf("%s%s", cacheWorkbenchJwtBlacklistsJwtPrefix, jwt)
	v := JwtBaseInfo{
		Uid:    uid,
		Status: status,
	}
	res, err := json.Marshal(v)
	if err != nil {
		return errorx.InternalError("服务器错误", err)
	}
	err = global.RedisX.Setex(workbenchJwtBlacklistsJwtKey, string(res), int(exp))
	if err != nil {
		return errorx.InternalError("服务器错误", err)
	}
	return nil
}
