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

	"workbench-server/common/captchax"
	"workbench-server/common/errorx"
	"workbench-server/common/global"
	"workbench-server/common/jwtx"
	"workbench-server/common/respx"
	"workbench-server/common/utils"
	"workbench-server/service/v1/app/api/internal/svc"
	"workbench-server/service/v1/app/api/internal/types"
	"workbench-server/service/v1/app/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"go.uber.org/zap"
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

/** Login 用户登录
 * ①获取登录的所有信息，验证码、账号、密码、ip与platform
 * ②验证码校验
 * ③ip登录次数校验
 * ④账号类型判断：邮箱？手机号？账号
 * ⑤账号密码校验：
 **** 1、如果校验失败：做登录失败记录（消息队列）
 **** 2、如果校验成功：更新最后一次登录时间，总登录次数+1（延迟队列）；设备单点登录，强制注销另一个相同设备的token（消息队列）
 * ⑥返回 jwt 给前端
 */
func (l *LoginLogic) Login(r *http.Request, req types.UserLoginReq) (any, error) {
	// ①获取登录的所有信息，验证码、账号、密码、ip与platform
	solution := strings.TrimSpace(req.Solution)
	if len(solution) == 0 {
		return nil, errors.New("请输入验证码")
	}

	captchaKey := strings.TrimSpace(req.CaptchaKey)
	if len(captchaKey) == 0 {
		return nil, errors.New("验证码ID丢失")
	}

	passport := strings.TrimSpace(req.Passport)
	if len(passport) == 0 {
		return nil, errors.New("账号不能为空")
	}
	password := strings.TrimSpace(req.Password)
	if len(password) == 0 {
		return nil, errors.New("密码不能为空")
	}

	reqIp := utils.RemoteIp(r)

	platform, err := utils.RemotePlatform(r)
	if err != nil {
		return nil, err
	}

	// ②验证码校验
	if !captchax.Store.Verify(captchaKey, solution, true) {
		return nil, errors.New("验证码错误")
	}

	// verifyResult := base64Captcha.DefaultMemStore.Verify(captchaKey, solution, true)
	// if !verifyResult {
	// 	return nil, errors.New("验证码错误")
	// }

	// ③ip登录次数校验
	failnum, err := l.svcCtx.LogModel.FindLoginFailCount(reqIp)
	if err != nil {
		global.ZAP.Error(fmt.Sprintf("获取IP[%s]在今天内的登录错误的次数失败", reqIp), zap.Error(err))
		return nil, errorx.InternalError("查询登录次数错误", err)
	}

	if failnum > 10 {
		global.ZAP.Info(fmt.Sprintf("ip:%s今日登录失败次数过多", reqIp))
		return nil, errors.New("sorry! 你今天登录失败次数太多了，明天再尝试")
	}

	// ④账号类型判断：邮箱？手机号？账号
	var userInfo *model.User
	if utils.VerifyEmailFormat(passport) == nil {
		if userInfo, err = l.svcCtx.UserModel.FindOneByEmail(passport); err == model.ErrNotFound {
			return nil, errors.New("此邮箱未注册")
		}
	} else if utils.VerifyMobileFormat(passport) == nil {
		if userInfo, err = l.svcCtx.UserModel.FindOneByMobile(passport); err == model.ErrNotFound {
			return nil, errors.New("此手机号未注册")
		}
	} else {
		if userInfo, err = l.svcCtx.UserModel.FindOneByPassport(passport); err == model.ErrNotFound {
			return nil, errors.New("此账号未注册")
		}
	}
	if err != nil {
		return nil, errorx.InternalError("服务器异常", err)
	}

	// 登录日志记录
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, errorx.InternalError("json.Marshal failed，Req转json失败", err)
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

	//明文转密文,使用md5加密，不可逆
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	// 如果校验失败：做登录失败记录（消息队列）
	if userInfo.Password != password {
		LogInfo.Message = "密码错误"
		if e := l.AddLoginLog(LogInfo); e != nil {
			return nil, errorx.InternalError("登录失败,日志错误", e)
		}
		global.ZAP.Error(fmt.Sprintf("%s使用%s登录密码错误", passport, password))
		return nil, errors.New("密码错误")
	}
	//如果校验成功
	Duration := jwtx.DefaultDuration //有效期 1 天
	if req.AutoLogin {
		Duration = Duration * 7 // 7天内免登录
	}
	jwtObj := jwtx.NewJWT()
	baseClaims := jwtx.BaseClaims{
		Uid:      userInfo.Uid,
		Ip:       reqIp,
		Platform: platform,
		Duration: Duration,
	}

	//构建jwt
	customClaims := jwtObj.CreateMapClaims(baseClaims)
	jwtToken, err := jwtObj.CreateToken(customClaims)
	if err != nil {
		return nil, errorx.InternalError("服务器异常", err)
	}

	Resp := &types.JwtToken{
		AccessToken:  jwtToken,
		AccessExpire: time.Now().Unix() + Duration,
	}
	RespStr, err := json.Marshal(Resp)
	if err != nil {
		return nil, errorx.InternalError("json.Marshal failed，Resp转json失败", err)
	}

	// 更新登录信息，注意：务必再添加登录成功日志之前执行
	if e := l.PutLoginInfo(userInfo.Uid); e != nil {
		return nil, errorx.InternalError("登录失败,更新登录信息失败", e)
	}

	//移除相同设备已登录token，注意：务必再添加登录成功日志之前执行
	//强制注销另一个相同设备的token（消息队列）
	l.LoginOutToken(userInfo.Uid, platform)
	if err := l.LoginOutToken(userInfo.Uid, platform); err != nil {
		return nil, err
	}

	//登录成功日志
	LogInfo.Message = "登录成功"
	LogInfo.Status = 1
	LogInfo.Resp = string(RespStr)
	LogInfo.JwtToken = jwtToken
	LogInfo.UserId = userInfo.Uid
	if e := l.AddLoginLog(LogInfo); e != nil {
		return nil, errorx.InternalError("登录失败,日志错误", e)
	}

	return respx.SingleResp("登录成功！", Resp), nil
}

// func (l *LoginLogic) CreateJwtToken(iat, seconds, userId, userRoleId, userCompanyId int64, ip, platform string) (string, error) {
// 	claims := make(jwt.MapClaims)
// 	claims["exp"] = iat + seconds       // 过期时间 24h  配置文件，go-zero这垃圾框架后不会解析这个值，最好自己重新配一个，为下：ExpiresAt
// 	claims["ExpiresAt"] = iat + seconds // 过期指定时间，单位s，自己额外配置的
// 	claims["iat"] = iat                 // 签名生效时间
// 	claims["userId"] = userId
// 	claims["userRoleId"] = userRoleId
// 	claims["userCompanyId"] = userCompanyId
// 	claims["ip"] = ip
// 	claims["platform"] = platform
// 	// 创建一个新的令牌对象，指定签名方法和声明
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	// 使用密码签名并获得完整的编码令牌作为字符串
// 	return token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
// }

func (l *LoginLogic) AddLoginLog(data *model.SysLog) error {
	if _, err := l.svcCtx.LogModel.Insert(data); err != nil {
		global.ZAP.Error(fmt.Sprintf("[%v]日志添加失败", *data), zap.Error(err))
		return err
	}
	return nil
}

//PutLoginInfo 更新最后一次登录时间与总登录次数，更新至用户信息
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
			return fmt.Errorf("查询最后登录日志出错:%s", err)
		}
	}

	return l.svcCtx.UserModel.UpdateLoginInfo(uid, LoginCount, LastLoginLog.CreateTime)
}

//LoginOutToken 平台单点登录-主动注销处理，登录当前平台后，会主动注销之前的token
func (l *LoginLogic) LoginOutToken(uid int64, platform string) error {
	//查询日志，寻找上一个相同登录平台的token
	lastLog, err := l.svcCtx.LogModel.FindLastLogin(uid, platform)
	if err == model.ErrNotFound {
		return nil
	}
	if err != nil {
		return errorx.InternalError("获取最后一次登录日志失败", err)
	}
	if len(lastLog.JwtToken) == 0 {
		return errorx.InternalError("jwtToken不存在", nil)
	}

	TokenKey := fmt.Sprintf("token:blacklist:%s", lastLog.JwtToken)

	//判断redis中是否存在已注销token
	val, err := l.svcCtx.Redis.Get(TokenKey)
	if err != nil {
		return errorx.InternalError("redis服务器错误", err)
	}
	if val == "logout" {
		//发现已经被主动注销了，无需重复注销
		return nil
	}

	// 从token中取过期时间
	jwtObj := jwtx.NewJWT()
	claims, err := jwtObj.ParseToken(lastLog.JwtToken)
	if err != nil {
		return errorx.InternalError("jwtToken解析失败", err)
	}

	if claims.Valid() != nil {
		return nil
	}

	exp := jwtx.ExpiredTime(*claims)
	now := time.Now().Unix()
	dura := exp - now
	//token本身已经过期了，不用销毁
	if dura <= 0 {
		return nil
	}
	//此token未过期，需要要销毁
	if err := l.svcCtx.Redis.Setex(TokenKey, "repeal", int(dura)); err != nil {
		return errorx.InternalError("token销毁失败", err)
	}
	return nil
}
