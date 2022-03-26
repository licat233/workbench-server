package captcha

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"
	"workbench-server/common/respx"
	"workbench-server/common/utils"

	"github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/core/logx"
)

type EmailCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) EmailCaptchaLogic {
	return EmailCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailCaptchaLogic) EmailCaptcha(req *types.EmailCaptchaReq) (interface{}, error) {
	if utils.VerifyEmailFormat(req.Address) != nil {
		return nil, errors.New("郵箱格式錯誤")
	}
	rand.Seed(time.Now().UnixNano())
	random_number := rand.Intn(899999) + 100000
	random_string := strconv.Itoa(random_number)
	err := SendEmail(req.Address, random_string)
	return respx.NewDefaultResponse(err, nil)
}

/*
SendEmail 发送邮件
@target 收件地址
@theme 邮件主题
@content 邮件内容
*/
func SendEmail(target, captcha string) error {
	targetAddr := []string{target}
	content := fmt.Sprintf("你正在使用该邮箱注册享购工作台账号，验证码为:%s，2分钟内有效。", captcha)
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "licat <licat233@qq.com>"
	// 设置接收方的邮箱
	e.To = targetAddr
	//设置秘密抄送
	e.Bcc = targetAddr
	//设置主题
	e.Subject = "享购工作台验证码"
	//设置文件发送的内容
	e.Text = []byte(content)
	//设置服务器相关的配置
	return e.Send("smtp.qq.com:25", smtp.PlainAuth("", "licat233@qq.com", "mtfhjtjtcanmdeah", "smtp.qq.com"))
}
