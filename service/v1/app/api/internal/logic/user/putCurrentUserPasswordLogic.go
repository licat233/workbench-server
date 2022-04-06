package user

import (
	"context"
	"crypto/md5"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"workbench-server/common/errorx"
	"workbench-server/common/global"
	"workbench-server/common/respx"
	"workbench-server/common/tool"
	"workbench-server/common/utils"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
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

func (l *PutCurrentUserPasswordLogic) PutCurrentUserPassword(req *types.UserPutPwdReq, r *http.Request) (any, error) {
	uid, err := tool.GetJwtUid(l.ctx)
	if err != nil {
		return nil, errorx.ExternalError("用户密码更改失败！", err)
	}
	//校验旧密码
	oldPassword := strings.TrimSpace(req.OldPassword)
	if err := utils.VerifyPasswordFormat(oldPassword); err != nil {
		return nil, errorx.ExternalError("旧密码错误", err)
	}
	if info, err := l.svcCtx.UserModel.FindOne(uid); err != nil {
		return nil, errorx.InternalError("更改失败", err)
	} else if info.Password != fmt.Sprintf("%x", md5.Sum([]byte(oldPassword))) {
		return nil, errorx.ExternalError("旧密码错误", err)
	}

	//校验并设置新密码
	newPassword := strings.TrimSpace(req.NewPassword)
	if err := utils.VerifyPasswordFormat(newPassword); err != nil {
		return nil, errorx.ExternalError(err.Error(), nil)
	}
	newPassword = fmt.Sprintf("%x", md5.Sum([]byte(newPassword)))
	if err = l.svcCtx.UserModel.UpdatePassword(uid, newPassword); err != nil {
		return respx.Msg(err, nil, "密码更改成功", "密码更改失败")
	}

	if err := LogoutHandler(l.ctx, r); err != nil {
		global.ZAP.Error("密码更改成功，注销token失败:%s", zap.Error(err))
		return nil, errorx.InternalError("密码更改成功，注销token失败", err)
	}
	return respx.StateResp("密码更改成功，当前登录已注销，请使用新密码重新登录"), nil
}

func (l *PutCurrentUserPasswordLogic) Logout(r *http.Request) (any, error) {
	exp, err := strconv.ParseInt(fmt.Sprintf("%v", l.ctx.Value("ExpiresAt")), 10, 64)
	if err != nil {
		return nil, errorx.InternalError("获取过期时间失败", err)
	}
	now := time.Now().Unix()
	dura := exp - now
	if dura <= 0 {
		return nil, errorx.InternalError("当前登录已过期，无需主动注销", nil)
	}
	jwtToken := r.Header.Get("Authorization")
	token := fmt.Sprintf("token:blacklist:%s", jwtToken)
	err = l.svcCtx.Redis.Setex(token, "logout", int(dura))
	return respx.Msg(err, nil, "注销成功", "注销失败:Redis错误")
}
