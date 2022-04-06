package tool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

func GetJwtUid(ctx context.Context) (int64, error) {
	jwtUserId := ctx.Value("uid")
	uid, err := json.Number(fmt.Sprintf("%v", jwtUserId)).Int64()
	if err != nil {
		return 0, errors.New("获取不到你的UID")
	}
	return uid, nil
}

// func GetJwtRoleId(ctx context.Context) (int64, error) {
// 	JwtRoleId := ctx.Value("userRoleId")
// 	uid, err := json.Number(fmt.Sprintf("%v", JwtRoleId)).Int64()
// 	if err != nil {
// 		return 0, errors.New("获取不到你角色ID")
// 	}
// 	return uid, nil
// }

// func GetJwtCompanyId(ctx context.Context) (int64, error) {
// 	JwtUserCompanyId := ctx.Value("userCompanyId")
// 	uid, err := json.Number(fmt.Sprintf("%v", JwtUserCompanyId)).Int64()
// 	if err != nil {
// 		return 0, errors.New("获取不到你所属公司ID")
// 	}
// 	return uid, nil
// }
