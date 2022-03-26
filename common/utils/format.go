package utils

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

//VerifyEmailFormat 邮箱验证
func VerifyEmailFormat(email string) error {
	email = strings.TrimSpace(email)
	if len(email) == 0 {
		return errors.New("邮箱不能为空")
	}
	if len(email) >= 40 {
		return errors.New("邮箱长度超过系统限制")
	}
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	if reg.MatchString(email) {
		return nil
	}
	return errors.New("邮箱格式错误")
}

//VerifyMobileFormat 手机号验证
func VerifyMobileFormat(mobileNum string) error {
	mobileNum = strings.TrimSpace(mobileNum)
	if len(mobileNum) == 0 {
		return errors.New("手机号不能为空")
	}
	if len(mobileNum) >= 20 {
		return errors.New("手机号长度超过系统限制")
	}
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	if reg.MatchString(mobileNum) {
		return nil
	}
	return errors.New("手机号格式错误")
}

//VerifyRealnameFormat 用户姓名验证，必须含有中文
func VerifyRealnameFormat(name string) error {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return errors.New("真实姓名不能为空")
	}
	if len(name) >= 20 {
		return errors.New("真实姓名长度超过系统限制")
	}
	reg := regexp.MustCompile("^[\u4e00-\u9fa5]$")
	for _, v := range name {
		if reg.MatchString(string(v)) {
			return nil
		}
	}
	return errors.New("真实姓名格式错误")
}

//VerifyPassportFormat 判断账号是否符合规定，不能含有@、纯数字、中文
func VerifyPassportFormat(passport string) error {
	passport = strings.TrimSpace(passport)
	if len(passport) == 0 {
		return errors.New("账号不能为空")
	}
	if len(passport) >= 30 {
		return errors.New("账号长度超过系统限制")
	}
	if VerifyMobileFormat(passport) == nil || VerifyEmailFormat(passport) == nil || VerifyRealnameFormat(passport) == nil {
		return errors.New("账号格式错误，不能含有@和中文，请检查修改后再提交")
	}
	return nil
}

//VerifyPasswordFormat 验证密码格式
func VerifyPasswordFormat(password string) error {
	if govalidator.IsMD5(password) {
		return errors.New("密码格式错误，请勿传入MD5密文")
	}
	password = strings.TrimSpace(password)
	if len(password) == 0 {
		return errors.New("密码不能为空")
	}
	if VerifyRealnameFormat(password) == nil {
		return errors.New("密码不能含有中文")
	}
	reg := regexp.MustCompile("^[0-9A-Za-z]{6,16}$")
	if reg.MatchString(password) {
		return nil
	}
	return errors.New("密码必须由6-16位数字或字母组成")
}

//VerifyUidFormat 验证是不是uid
func VerifyUidFormat(uid string) error {
	uid = strings.TrimSpace(uid)
	if len(uid) < 6 && govalidator.IsInt(uid) { //用户id不可能大于 99999 的
		return nil
	}
	return errors.New("不是用户ID")
}

//StringToTimeString 时间字符串处理，最终返回”2006-01-02 00:00:00"格式，用于model
func StringToTimeString(t string) (string, error) {
	stamp, err := time.ParseInLocation("2006-01-02", t, time.Local)
	if err != nil {
		return "", errors.New("请传入正确时间字符串格式，如：2006-01-02")
	}
	return stamp.Format("2006-01-02 15:04:05"), nil
}

//StringToDate 时间字符串处理，最终返回日期时间
func StringToDate(t string) (time.Time, error) {
	stamp, err := time.ParseInLocation("2006-01-02", t, time.Local)
	if err != nil {
		return time.Time{}, errors.New("请传入正确时间字符串格式，如：2006-01-02")
	}
	return stamp, nil
}

//VerifyIpFormat ip地址验证
func VerifyIpFormat(ip string) error {
	reg := regexp.MustCompile(`^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`)
	if reg.MatchString(ip) {
		return nil
	}
	return errors.New("获取不到客户端ip地址")
}
