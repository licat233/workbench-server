package user

import (
	"context"
	"crypto/md5"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"workbench-server/global/errorx"
	"workbench-server/global/respx"
	"workbench-server/global/utils"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutCurrentUserPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutCurrentUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) PutCurrentUserPasswordLogic {
	return PutCurrentUserPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutCurrentUserPasswordLogic) PutCurrentUserPassword(req types.UserPutPwdReq, r *http.Request) (interface{}, error) {
	uid, err := utils.GetJwtUid(l.ctx)
	if err != nil {
		return nil, errorx.NewRequestError("用户密码更改失败！", err)
	}
	//校验旧密码
	oldPassword := strings.TrimSpace(req.OldPassword)
	if err := utils.VerifyPasswordFormat(oldPassword); err != nil {
		return nil, errorx.NewRequestError("旧密码错误", err)
	}
	if info, err := l.svcCtx.UserModel.FindOne(uid); err != nil {
		return nil, errorx.NewServerError("更改失败", err)
	} else if info.Password != fmt.Sprintf("%x", md5.Sum([]byte(oldPassword))) {
		return nil, errorx.NewRequestError("旧密码错误", err)
	}

	//校验并设置新密码
	newPassword := strings.TrimSpace(req.NewPassword)
	if err := utils.VerifyPasswordFormat(newPassword); err != nil {
		return nil, errorx.NewRequestError(err.Error(), nil)
	}
	newPassword = fmt.Sprintf("%x", md5.Sum([]byte(newPassword)))
	if err = l.svcCtx.UserModel.UpdatePassword(uid, newPassword); err != nil {
		return respx.NewDefaultResponse(err, nil, "密码更改成功", "密码更改失败")
	}
	if _, err := l.Logout(r); err != nil {
		logx.Errorf("密码更改成功，注销当前token失败:%s", err)
	}
	return respx.NewDefaultResponse(err, nil, "密码更改成功，当前登录已注销，请使用新密码重新登录", "密码更改成功，注销当前token失败")
}

func (l *PutCurrentUserPasswordLogic) Logout(r *http.Request) (interface{}, error) {
	exp, err := strconv.ParseInt(fmt.Sprintf("%v", l.ctx.Value("ExpiresAt")), 10, 64)
	if err != nil {
		return nil, errorx.NewDefaultError("获取过期时间失败", err)
	}
	now := time.Now().Unix()
	dura := exp - now
	if dura <= 0 {
		return nil, errorx.NewDefaultError("当前登录已过期，无需主动注销", nil)
	}
	jwtToken := r.Header.Get("Authorization")
	token := fmt.Sprintf("token:blacklist:%s", jwtToken)
	err = l.svcCtx.RedisClient.Setex(token, "logout", int(dura))
	return respx.NewDefaultResponse(err, nil, "注销成功", "注销失败:Redis错误")
}
