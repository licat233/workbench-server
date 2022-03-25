package errorx

import (
	"fmt"
	"net/http"
)

const InternalCode = 500
const ExternalCode = 400
const TraceMode = true //调试模式打开
const ShowType = 2

/** 与前端约定的数据规范
interface ErrorInfoStructure {
	success: boolean; // 请求状态
	data?: any; // 返回数据
	errorCode?: string; // 错误类型代码：400错误请求，401未授权，500服务器内部错误，200成功
	errorMessage?: string; // 向用户显示消息
	traceMessage?: string; // 调试错误信息，请勿在生产环境下使用，可有可无
	showType?: number; // 错误显示类型：0无提示；1.信息警告；2.信息错误；4通知；9跳转到指定页
	traceId?: string; // 方便后端故障排除：唯一的请求ID
	host?: string; // 方便后端故障排除：当前访问服务器的主机
}

export enum ErrorShowType {
  SILENT = 0, // 不提示错误
  WARN_MESSAGE = 1, // 警告信息提示
  ERROR_MESSAGE = 2, // 错误信息提示
  NOTIFICATION = 4, // 通知提示
  REDIRECT = 9, // 页面跳转
}
**/
type CodeError struct {
	Status bool   `json:"success"`
	Code   int    `json:"errorCode"`
	Msg    string `json:"errorMessage"`
	Type   int    `json:"showType"`
	Trace  string `json:"traceMessage"`
}

//用于返回的错误结构体
type CodeErrorResponse struct {
	Status bool   `json:"success"`
	Code   int    `json:"errorCode"`
	Msg    string `json:"errorMessage"`
	Type   int    `json:"showType"`
	Trace  string `json:"traceMessage"`
}

//NewCodeError 新建一个自定义错误实例
func NewCodeError(code int, msg string, trace error) error {
	var traceMsg string
	if TraceMode && trace != nil {
		traceMsg = trace.Error()
	}
	return &CodeError{
		Code:  code,
		Msg:   msg,
		Type:  ShowType,
		Trace: traceMsg,
	}
}

//NewDefaultError 默认为服务器内部错误 500
func NewDefaultError(msg string, trace error) error {
	return NewServerError(msg, trace)
}

//NewServerError 服务器错误 500
func NewServerError(msg string, trace error) error {
	var traceMsg string
	if TraceMode && trace != nil {
		traceMsg = trace.Error()
	}
	return &CodeError{
		Code:  InternalCode,
		Msg:   msg,
		Type:  2,
		Trace: traceMsg,
	}
}

//NewRequestError 请求错误 400
func NewRequestError(msg string, trace error) error {
	var traceMsg string
	if TraceMode && trace != nil {
		traceMsg = trace.Error()
	}
	return &CodeError{
		Code:  ExternalCode,
		Msg:   msg,
		Type:  1,
		Trace: traceMsg,
	}
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:  ExternalCode,
		Msg:   e.Msg,
		Type:  ShowType,
		Trace: e.Trace,
	}
}

//NewDefaultError 为请求错误 400
func NewErrorResponse(msg string, trace error) *CodeErrorResponse {
	var traceMsg string
	if TraceMode && trace != nil {
		traceMsg = trace.Error()
	}
	return &CodeErrorResponse{
		Code:  ExternalCode,
		Msg:   msg,
		Type:  ShowType,
		Trace: traceMsg,
	}
}

func (e *CodeError) Error() string {
	return e.Msg
}

//使用 NewDefaultError 建立的错误，为服务器内部错误
func ErrorHandler(err error) (int, interface{}) {
	if err == nil {
		return http.StatusOK, NewCodeError(200, "", nil)
	}
	//一般服务器内部有问题才会用到这个错误
	if v, b := err.(*CodeError); b {
		return http.StatusOK, v.Data()
	}
	// 设置http错误400或者500会导致前端直接获取不到报错信息，只能设置为 200
	// return http.StatusBadRequest, NewErrorResponse(fmt.Sprint(err), nil)
	return http.StatusOK, NewErrorResponse(fmt.Sprint(err), nil)
}
