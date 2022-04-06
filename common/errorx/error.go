package errorx

import (
	"errors"
	"net/http"
	"workbench-server/common/global"

	"go.uber.org/zap"
)

/** Errorx 对接前端antpro框架的接口格式规范
 */
type Errorx struct {
	Success      bool   `json:"success"`      // 【必需】请求状态
	ErrorCode    int    `json:"errorCode"`    // 【选填】错误类型代码：400错误请求，401未授权，500服务器内部错误，200成功
	ErrorMessage string `json:"errorMessage"` // 【选填】向用户显示消息
	TraceMessage string `json:"traceMessage"` // 【选填】调试错误信息，请勿在生产环境下使用，可有可无
	ShowType     int    `json:"showType"`     // 【选填】错误显示类型：0.不提示错误;1.警告信息提示；2.错误信息提示；4.通知提示；9.页面跳转
	TraceId      string `json:"traceId"`      // 【选填】方便后端故障排除：唯一的请求ID
	Host         string `json:"host"`         // 【选填】方便后端故障排除：当前访问服务器的主机
}

const (
	ERROR_CODE_INTERNAL = 500  //内部错误
	ERROR_CODE_EXTERNAL = 400  //外部错误（默认）
	ERROR_CODE_AUTH     = 401  //外部错误（默认）
	SHOW_TYPE           = 2    //信息错误
	TRACE_MODEL         = true //调试模式
)

//Error 返回错误信息，主要用于实现原生error接口
func (e *Errorx) Error() string {
	return e.ErrorMessage
}

//newErrorx 新建Errorx实例
func newErrorx(alertMsg string, traceErr error) *Errorx {
	traceMsg := ""
	if TRACE_MODEL {
		traceMsg = traceErr.Error()
	}
	return &Errorx{
		Success:      false,
		ErrorCode:    ERROR_CODE_EXTERNAL,
		ErrorMessage: alertMsg,
		TraceMessage: traceMsg,
		ShowType:     SHOW_TYPE,
		TraceId:      "",
		Host:         "",
	}
}

/** New 新建Errorx错误，默认是EXTERNAL错误，因为INTERNAL错误比较少
 * 需要传入以下参数，分别为:
 * errMsg string 提示信息
 * traceErr error 调试信息
 * errCode int 错误代码
 * showType int 错误显示类型,0.不提示错误;1.警告信息提示；2.错误信息提示；4.通知提示；9.页面跳转
 */
func New(params ...any) error {
	var traceErr error
	errMsg, errCode, showtype := "请求失败", ERROR_CODE_EXTERNAL, SHOW_TYPE
	n := len(params)
	if n == 1 {
		p := params[0]
		if p != nil {
			if v, ok := p.(string); ok {
				if len(v) > 0 {
					errMsg = v
				}
				goto OVER
			}
			if v, ok := p.(error); ok {
				traceErr = v
				goto OVER
			}
			if v, ok := p.(int); ok {
				if v > 10 {
					errCode = v
				} else {
					showtype = v
				}
				goto OVER
			}
		}
		goto OVER
	}
	if n > 1 {
		if params[0] != nil {
			if v, ok := params[0].(string); ok {
				errMsg = v
			} else {
				return InternalError("服务器错误", errors.New("errorx.New()的第1个参数不是string"))
			}
		}
		if params[1] != nil {
			if v, ok := params[1].(error); ok {
				traceErr = v
			} else {
				return InternalError("服务器错误", errors.New("errorx.New()的第2个参数不是error"))
			}
		}
	}
	if n > 2 {
		if params[2] != nil {
			if v, ok := params[2].(int); ok {
				errCode = v
			} else {
				return InternalError("服务器错误", errors.New("errorx.New()的第3个参数不是int"))
			}
		}
	}
	if n > 3 {
		if params[3] != nil {
			if v, ok := params[3].(int); ok {
				showtype = v
			} else {
				return InternalError("服务器错误", errors.New("errorx.New()的第4个参数不是int"))
			}
		}
	}
OVER:
	err := newErrorx(errMsg, traceErr)
	err.ErrorCode = errCode
	err.ShowType = showtype
	return err
}

/** InternalError 内部错误 500
 * alertMsg string 提示信息
 * traceErr string 调试信息
 */
func InternalError(alertMsg string, traceErr error) error {
	res := newErrorx(alertMsg, traceErr)
	if len(res.ErrorMessage) == 0 {
		res.ErrorMessage = "服务器错误"
	}
	res.ErrorCode = ERROR_CODE_INTERNAL
	//非调试模式，统一报：服务器错误
	if !TRACE_MODEL {
		res.ErrorMessage = "服务器错误"
	}
	//内部错误要做日志处理
	global.ZAP.Error("服务器错误", zap.Error(traceErr))
	return res
}

/** ServerError 服务器错误 500
 * alertMsg string 提示信息
 * traceMsg string 调试信息
 */
func ServerError(alertMsg string, traceErr error) error {
	return InternalError(alertMsg, traceErr)
}

/** ExternalError 外部错误 400
 * alertMsg string 提示信息
 * traceMsg string 调试信息
 */
func ExternalError(alertMsg string, traceErr error) error {
	res := newErrorx(alertMsg, traceErr)
	if len(res.ErrorMessage) == 0 {
		res.ErrorMessage = "请求错误"
	}
	res.ErrorCode = ERROR_CODE_EXTERNAL
	return res
}

/** RequestError 请求错误 400
 * alertMsg string 提示信息
 * traceMsg string 调试信息
 */
func RequestError(alertMsg string, traceErr error) error {
	return ExternalError(alertMsg, traceErr)
}

/** AuthError 身份校验失败 401
 * alertMsg string 提示信息
 * traceMsg string 调试信息
 */
func AuthError(alertMsg string, traceErr error) error {
	res := newErrorx(alertMsg, traceErr)
	if len(res.ErrorMessage) == 0 {
		res.ErrorMessage = "身份校验失败"
	}
	res.ErrorCode = ERROR_CODE_AUTH
	return res
}

/** ErrorToErrorX 对接go-zero的错误处理函数，将原生错误转化为 Errorx，默认为ExternalError
 * inerr 传入的错误
 */
func ErrorToErrorX(inerr error) (int, interface{}) {
	if inerr == nil {
		//没有错误，应该正常返回
		return http.StatusOK, &struct {
			Success bool `json:"success"` // 请求状态。
		}{
			Success: true,
		}
	}
	//类型断言看是否为 Errorx 类型
	if v, ok := inerr.(*Errorx); ok {
		//无需处理，直接返回
		return http.StatusOK, v
	}
	//原生错误，我们默认把原生error作为ExternalError
	return http.StatusOK, ExternalError(inerr.Error(), nil)
}
