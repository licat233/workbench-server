package user

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"workbench-server/app/api/v1/internal/svc"
	"workbench-server/app/api/v1/internal/types"
	"workbench-server/app/model"
	"workbench-server/common/errorx"
	"workbench-server/common/respx"
	"workbench-server/common/utils"

	"github.com/golang-jwt/jwt"
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(r *http.Request, req types.UserLoginReq) (interface{}, error) {
	solution := strings.TrimSpace(req.Solution)
	if len(solution) == 0 {
		return nil, errorx.NewRequestError("请输入验证码", nil)
	}

	CaptchaKey := strings.TrimSpace(req.CaptchaKey)
	if len(CaptchaKey) == 0 {
		return nil, errorx.NewRequestError("验证码ID丢失", nil)
	}

	passport := strings.TrimSpace(req.Passport)
	password := strings.TrimSpace(req.Password)
	if len(passport) == 0 {
		return nil, errorx.NewRequestError("账号不能为空", nil)
	}
	if len(password) == 0 {
		return nil, errorx.NewRequestError("密码不能为空", nil)
	}

	verifyResult := base64Captcha.DefaultMemStore.Verify(req.CaptchaKey, req.Solution, true)

	if !verifyResult {
		return nil, errorx.NewRequestError("验证码错误", nil)
	}

	reqIp := utils.RemoteIp(r)
	failnum, err := l.svcCtx.LogModel.FindLoginFailCount(reqIp)
	if err != nil {
		logx.Errorf("获取IP[%s]在今天内的登录错误的次数失败:%s", reqIp, err)
		return nil, errorx.NewServerError("查询登录次数错误", err)
	}
	if failnum > 10 {
		return nil, errorx.NewRequestError("sorry! 你今天登录失败次数太多了，明天再尝试", nil)
	}

	var userInfo *model.User
	if utils.VerifyEmailFormat(passport) == nil {
		if userInfo, err = l.svcCtx.UserModel.FindOneByEmail(passport); err == model.ErrNotFound {
			return nil, errorx.NewRequestError("此邮箱未注册", nil)
		}
	} else if utils.VerifyMobileFormat(passport) == nil {
		if userInfo, err = l.svcCtx.UserModel.FindOneByMobile(passport); err == model.ErrNotFound {
			return nil, errorx.NewRequestError("此手机号未注册", nil)
		}
	} else {
		if userInfo, err = l.svcCtx.UserModel.FindOneByPassport(passport); err == model.ErrNotFound {
			return nil, errorx.NewRequestError("此账号未注册", nil)
		}
	}
	if err != nil {
		return nil, errorx.NewServerError("服务器异常", err)
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, errorx.NewServerError("json.Marshal failed，Req转json失败", err)
	}

	platform, err := utils.RemotePlatform(r)
	if err != nil {
		return nil, err
	}

	LogInfo := &model.SysLog{
		Ip:       reqIp,
		Method:   "post",
		Path:     "/user/login",
		Status:   0,
		Message:  "",
		Body:     string(reqBody),
		Resp:     "",
		UserId:   0,
		JwtToken: "",
		Platform: platform,
	}

	if err != nil {
		LogInfo.Message = fmt.Sprintf("登录失败，%s", err)
		if e := l.AddLoginLog(LogInfo); e != nil {
			return nil, errorx.NewDefaultError("登录失败,日志错误", e)
		}
		return nil, errorx.NewDefaultError("登录失败", err)
	}

	//明文转密文,使用md5加密，不可逆
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	if userInfo.Password != password {
		LogInfo.Message = "密码错误"
		if e := l.AddLoginLog(LogInfo); e != nil {
			return nil, errorx.NewDefaultError("登录失败,日志错误", e)
		}
		logx.Error(fmt.Sprintf("%s使用%s登录密码错误", passport, password))
		return nil, errorx.NewRequestError("密码错误", nil)
	}

	accessExpire := l.svcCtx.Config.Auth.AccessExpire //有效期 1 天
	if req.AutoLogin {
		accessExpire = accessExpire * 7 // 7天内免登录
	}
	now := time.Now().Local().Unix()
	jwtToken, err := l.CreateJwtToken(now, accessExpire, userInfo.Uid, userInfo.UserRoleId, userInfo.UserCompanyId, reqIp, platform)
	if err != nil {
		logx.Errorf("创建JwtToken出错:%s", err)
		return nil, errorx.NewDefaultError("创建JwtToken出错", err)
	}
	Resp := &types.JwtToken{
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
	}
	RespStr, err := json.Marshal(Resp)
	if err != nil {
		return nil, errorx.NewDefaultError("json.Marshal failed，Resp转json失败", err)
	}

	// 更新登录信息，注意：务必再添加登录成功日志之前执行
	if e := l.PutLoginInfo(userInfo.Uid); e != nil {
		return nil, errorx.NewDefaultError("登录失败,更新登录信息失败", e)
	}

	//移除相同设备已登录token，注意：务必再添加登录成功日志之前执行
	if err := l.LoginOutToken(userInfo.Uid, platform); err != nil {
		logx.Error(err)
	}

	//登录成功日志
	LogInfo.Message = "登录成功"
	LogInfo.Status = 1
	LogInfo.Resp = string(RespStr)
	LogInfo.JwtToken = jwtToken
	LogInfo.UserId = userInfo.Uid
	if e := l.AddLoginLog(LogInfo); e != nil {
		return nil, errorx.NewDefaultError("登录失败,日志错误", e)
	}

	return respx.NewResponse("登录成功！", Resp), nil
}

func (l *LoginLogic) CreateJwtToken(iat, seconds, userId, userRoleId, userCompanyId int64, ip, platform string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds       // 过期时间 24h  配置文件，go-zero这垃圾框架后不会解析这个值，最好自己重新配一个，这里配置为：ExpiresAt
	claims["iat"] = iat                 // 签名生效时间
	claims["ExpiresAt"] = iat + seconds // 过期指定时间，单位s，自己额外配置的
	claims["userId"] = userId
	claims["userRoleId"] = userRoleId
	claims["userCompanyId"] = userCompanyId
	claims["ip"] = ip
	claims["platform"] = platform
	// 创建一个新的令牌对象，指定签名方法和声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密码签名并获得完整的编码令牌作为字符串
	return token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
}

func (l *LoginLogic) AddLoginLog(data *model.SysLog) error {
	if _, err := l.svcCtx.LogModel.Insert(data); err != nil {
		logx.Errorf("[%s]日志添加失败:%s", data, err)
		return err
	}
	return nil
}

//PutLoginInfo 获取最后一次登录时间与总登录次数，更新至用户信息
func (l *LoginLogic) PutLoginInfo(uid int64) error {
	LoginCount, err := l.svcCtx.LogModel.FindLoginCountByUid(uid)
	if err != nil {
		return fmt.Errorf("查询登录次数出错:%s", err)
	}

	//算上本次，所以 +1
	LoginCount++
	LastLoginLog, err := l.svcCtx.LogModel.FindLastLoginByUid(uid)
	if err != nil {
		if err == sqlc.ErrNotFound {
			//从未登录过，设为当前时间
			LastLoginLog.CreateTime = time.Now()
		} else {
			logx.Error(err)
			return fmt.Errorf("查询最后登录日志出错:%s", err)
		}
	}

	return l.svcCtx.UserModel.UpdateLoginInfo(uid, LoginCount, LastLoginLog.CreateTime)
}

//LoginOutToken 平台单点登录-主动注销处理
func (l *LoginLogic) LoginOutToken(uid int64, platform string) error {
	//查询日志，寻找上一个相同登录平台的token
	lastLog, err := l.svcCtx.LogModel.FindLastLogin(uid, platform)
	if err != nil {
		logx.Error(fmt.Errorf("查询登录日志出错! %s", err))
		return err
	}
	if len(lastLog.JwtToken) == 0 {
		return errors.New("jwtToken不存在")
	}

	TokenKey := fmt.Sprintf("token:blacklist:%s", lastLog.JwtToken)

	//判断redis中是否存在已注销token
	val, _ := l.svcCtx.RedisClient.Get(TokenKey)
	if val == "logout" {
		//发现已经被主动注销了，无需重复注销
		return nil
	}

	// 从token中取过期时间
	claims, err := l.svcCtx.JwtObject.ParseToken(lastLog.JwtToken)
	if err != nil {
		errmsg := fmt.Errorf("jwtToken解析失败 %s", err)
		logx.Error(errmsg)
		return errmsg
	}
	expstr := l.svcCtx.JwtObject.GetIdFromClaims("ExpiresAt", *claims)
	exp, err := utils.ExpToNumber(expstr)
	if err != nil {
		errmsg := fmt.Errorf("获取jwtToken过期时间失败 %s", err)
		return errmsg
	}

	now := time.Now().Unix()
	dura := exp - now
	//token本身已经过期了，不用销毁
	if dura < 0 {
		return nil
	}
	//此token未过期，需要要销毁
	if err := l.svcCtx.RedisClient.Setex(TokenKey, "repeal", int(dura)); err != nil {
		errmsg := fmt.Errorf("token销毁失败! %s", err)
		logx.Error(errmsg)
		return errmsg
	}
	return nil
}
