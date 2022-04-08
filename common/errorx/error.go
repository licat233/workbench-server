package errorx

import (
	"errors"
	"net/http"
	"workbench-server/common/global"

	"go.uber.org/zap"
)

/** Errorx 对接前端antpro框架的接口格式规范，用于实现error接口
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

/** ResponseError 对接前端antpro框架的接口格式规范，用于实现展现错误
 */
type ResponseError struct {
	Success      bool   `json:"success"`      // 【必需】请求状态
	ErrorCode    int    `json:"errorCode"`    // 【选填】错误类型代码：400错误请求，401未授权，500服务器内部错误，200成功
	ErrorMessage string `json:"errorMessage"` // 【选填】向用户显示消息
	TraceMessage string `json:"traceMessage"` // 【选填】调试错误信息，请勿在生产环境下使用，可有可无
	ShowType     int    `json:"showType"`     // 【选填】错误显示类型：0.不提示错误;1.警告信息提示；2.错误信息提示；4.通知提示；9.页面跳转
	TraceId      string `json:"traceId"`      // 【选填】方便后端故障排除：唯一的请求ID
	Host         string `json:"host"`         // 【选填】方便后端故障排除：当前访问服务器的主机
}

const (
	ERROR_CODE_INTERNAL = 500 //内部错误
	ERROR_CODE_EXTERNAL = 400 //外部错误（默认）
	ERROR_CODE_AUTH     = 401 //外部错误（默认）
	SHOW_TYPE           = 2   //信息错误
)

// var TRACE_MODEL = isTraceModel() //调试模式
var ServerError = InternalError
var RequestError = ExternalError

//Error 返回错误信息，主要用于实现原生error接口，
//注意：打印errorx类型的时候，会调用该函数值，所以返回errorx，等于返回该Error()的结果
func (e *Errorx) Error() string {
	return e.ErrorMessage
}

//Data 返回errorx的值
//如果需要返回errorx的值怎么办？
//那就需要我们自己定义一个展现errorx结果的函数，使用ResponseError结构体
func (e *Errorx) Data() *ResponseError {
	return &ResponseError{
		Success:      e.Success,
		ErrorCode:    e.ErrorCode,
		ErrorMessage: e.ErrorMessage,
		TraceMessage: e.TraceMessage,
		ShowType:     e.ShowType,
		TraceId:      e.TraceId,
		Host:         e.Host,
	}
}

/** InternalError 内部错误 500
 * alertMsg string 提示信息
 * traceErr string 调试信息
 */
func InternalError(alertMsg string, traceErr error) error {
	return internalError(alertMsg, traceErr)
}

/** ExternalError 外部错误 400
 * alertMsg string 提示信息
 * traceMsg string 调试信息
 */
func ExternalError(alertMsg string, traceErr error) error {
	return externalError(alertMsg, traceErr)
}

/** AuthError 身份校验失败 401
 * alertMsg string 提示信息
 * traceMsg string 调试信息
 */
func AuthError(alertMsg string, traceErr error) error {
	return authError(alertMsg, traceErr)
}

func internalError(alertMsg string, traceErr error) *Errorx {
	errorx := newErrorx(alertMsg, traceErr)
	if len(errorx.ErrorMessage) == 0 {
		errorx.ErrorMessage = "服务器错误"
	}
	errorx.ErrorCode = ERROR_CODE_INTERNAL
	//非调试模式，统一报：服务器错误
	if !isTraceModel() {
		errorx.ErrorMessage = "服务器错误"
		errorx.TraceMessage = "服务器出现错误，请联系管理员licat233@gmail.com"
	}
	//内部错误要做日志处理
	global.ZAP.Error("服务器错误", zap.Error(traceErr))
	return errorx
}

func externalError(alertMsg string, traceErr error) *Errorx {
	errorx := newErrorx(alertMsg, traceErr)
	if len(errorx.ErrorMessage) == 0 {
		errorx.ErrorMessage = "请求错误"
	}
	errorx.ErrorCode = ERROR_CODE_EXTERNAL
	return errorx
}

func authError(alertMsg string, traceErr error) *Errorx {
	errorx := newErrorx(alertMsg, traceErr)
	if len(errorx.ErrorMessage) == 0 {
		errorx.ErrorMessage = "身份校验失败"
	}
	errorx.ErrorCode = ERROR_CODE_AUTH
	return errorx
}

//newErrorx 新建Errorx实例，默认为外部请求错误
func newErrorx(alertMsg string, traceErr error) *Errorx {
	traceMsg := ""
	//调试模式下，展现错误调试信息
	if isTraceModel() && traceErr != nil {
		traceMsg = traceErr.Error()
	}
	errorx := &Errorx{
		Success:      false,
		ErrorCode:    ERROR_CODE_EXTERNAL,
		ErrorMessage: alertMsg,
		TraceMessage: traceMsg,
		ShowType:     SHOW_TYPE,
		TraceId:      "",
		Host:         "",
	}
	return errorx
}

/** ErrorToErrorX 对接go-zero的错误处理函数，将原生错误转化为 Errorx，默认为ExternalError
 * inerr 传入的错误
 */
func ErrorToErrorX(inerr error) (int, interface{}) {
	if inerr == nil {
		//没有错误，应该正常返回
		return http.StatusOK, &struct {
			Success bool   `json:"success"` // 请求状态。
			Message string `json:"message"` // 响应消息
		}{
			Success: true,
			Message: "请求成功",
		}
	}
	//类型断言看是否为 Errorx 类型
	if errorx, ok := inerr.(*Errorx); ok {
		//无需处理，直接返回
		return http.StatusOK, errorx.Data()
	}

	//原生错误，我们默认把原生error作为ExternalError返回
	// return http.StatusOK, externalError(inerr.Error(), nil)
	return http.StatusOK, externalError(inerr.Error(), nil).Data()
}

/** New 新建Errorx错误，默认是EXTERNAL错误，因为INTERNAL错误比较少
 * description：自定义语法糖，尽量少用，影响性能
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

func isTraceModel() bool {
	if global.CONFIG.ServerConfig.Mode == "Dev" || global.CONFIG.ServerConfig.Mode == "dev" {
		return true
	}
	return false
}
