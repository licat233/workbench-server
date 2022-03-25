package user

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"strings"

	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/common/utils"
	"workbench-server/service/app/api/v1/internal/svc"
	"workbench-server/service/app/api/v1/internal/types"
	"workbench-server/service/app/model"

	"github.com/mojocn/base64Captcha"
	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.UserRegisterReq) (interface{}, error) {
	solution := strings.TrimSpace(req.Solution)
	if len(solution) == 0 {
		return nil, errors.New("请输入验证码")
	}

	captchaKey := strings.TrimSpace(req.CaptchaKey)
	if len(captchaKey) == 0 {
		return nil, errors.New("验证码Key丢失")
	}
	ok := base64Captcha.DefaultMemStore.Verify(captchaKey, solution, true)
	if !ok {
		return nil, errors.New("验证码错误")
	}

	passport := strings.TrimSpace(req.Passport) //唯一
	realname := strings.TrimSpace(req.Name)
	email := strings.TrimSpace(req.Email)   //唯一
	mobile := strings.TrimSpace(req.Mobile) //唯一
	qq := strings.TrimSpace(req.QQ)
	wechat := strings.TrimSpace(req.Wechat)
	password := strings.TrimSpace(req.Password)

	if err := utils.VerifyRealnameFormat(realname); err != nil {
		return nil, err
	}
	//如果必须传入明文，才能校验
	if err := utils.VerifyPassportFormat(passport); err != nil {
		return nil, err
	}

	if err := utils.VerifyMobileFormat(mobile); err != nil {
		return nil, err
	}

	if err := utils.VerifyEmailFormat(email); err != nil {
		return nil, err
	}

	if err := utils.VerifyPasswordFormat(password); err != nil {
		return nil, err
	}

	if req.Gender != 0 && req.Gender != 1 {
		return nil, errors.New("'user_gender' parameter error")
	}

	if has, err := l.svcCtx.UserModel.IsHasByPassport(passport); err != nil {
		return nil, errorx.NewServerError("服务器出错", err)
	} else if has {
		return nil, errors.New("此账号已经被注册")
	}

	if has, err := l.svcCtx.UserModel.IsHasByMobile(mobile); err != nil {
		return nil, errorx.NewServerError("服务器出错", err)
	} else if has {
		return nil, errors.New("此手机号已经被注册")
	}

	if has, err := l.svcCtx.UserModel.IsHasByEmail(email); err != nil {
		return nil, errorx.NewServerError("服务器出错", err)
	} else if has {
		return nil, errors.New("此邮箱已经被注册")
	}

	//md5加密password，必须在校验完成以后才能md5加密
	password = fmt.Sprintf("%x", md5.Sum([]byte(password))) //MD5加密32位

	err := l.svcCtx.UserModel.InsertUserRegister(&model.User{
		Passport:      passport,
		UserRealName:  realname,
		UserAvatar:    "https://img.tukuppt.com/ad_preview/00/09/06/hyVORwxNGY.jpg",
		UserMobile:    mobile,
		UserEmail:     email,
		UserGender:    req.Gender,
		UserRoleId:    req.RoleId,
		UserCompanyId: req.CompanyId,
		UserQq:        qq,
		UserWechat:    wechat,
		Password:      password,
	})

	return respx.NewDefaultResponse(err, nil, "注册成功", "注册失败")
}
