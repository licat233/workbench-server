package respx

import (
	"workbench-server/app/model"
	"workbench-server/common/errorx"
	"workbench-server/common/utils"
)

type Result struct {
	Status  bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListResult struct {
	Status  bool        `json:"success"`
	Message string      `json:"message"`
	Page    int64       `json:"page"`
	Total   int64       `json:"total"`
	Data    interface{} `json:"data"`
}

func NewResponse(msg string, data interface{}) interface{} {
	return &Result{
		Status:  true,
		Message: msg,
		Data:    data,
	}
}

/*
NewDefaultResponse 应答规范:
{
	success: boolean
	message: string
	data: any
}
*/
func NewDefaultResponse(err error, data interface{}, msg ...string) (interface{}, error) {
	success := "请求成功"
	fail := "请求失败"
	for k, v := range msg {
		if k == 0 {
			success = v
		} else if k == 1 {
			fail = v
		}
	}
	if err == model.ErrNotFound {
		return NewResponse(success, nil), nil
	}
	if err != nil {
		return nil, errorx.NewDefaultError(fail, err)
	}
	return NewResponse(success, data), nil
}

// 用于返回list数据
func ListResponse(msg string, data interface{}, totalRecord, pageSize int64) interface{} {
	if !utils.IsSlice(data) {
		// return NewResponse(msg, data)
		data = []interface{}{data}
	}
	pageCount := (totalRecord + pageSize - 1) / pageSize
	return &ListResult{
		Status:  true,
		Message: msg,
		Page:    pageCount,
		Total:   totalRecord,
		Data:    data,
	}
}

func NewListResponse(err error, data interface{}, totalRecord, pageSize int64, msg ...string) (interface{}, error) {
	success := "请求成功"
	fail := "请求失败"
	for k, v := range msg {
		if k == 0 {
			success = v
		} else if k == 1 {
			fail = v
		}
	}
	if err == model.ErrNotFound {
		return NewResponse(success, []interface{}{}), nil
	}
	if err != nil {
		return nil, errorx.NewDefaultError(fail, err)
	}
	return ListResponse(success, data, totalRecord, pageSize), nil
}

/** 与前端约定的数据规范
interface ErrorInfoStructure {
	status: boolean; // 请求状态。可有可无
	success: boolean; // 请求状态
	page?: number; // 总共有多少页，根据前端的pageSize来计算
	total?: number; // 数据总个数
	data?: any; // 返回数据
	errorCode?: string; // 错误类型代码：400错误请求，401未授权，500服务器内部错误，200成功
	errorMessage?: string; // 向用户显示消息
	traceMessage?: string; // 调试错误信息，请勿在生产环境下使用
	showType?: number; // 错误显示类型：0无提示；1.信息警告；2.信息错误；4通知；9页
	traceId?: string; // 方便后端故障排除：唯一的请求ID
	host?: string; // 方便后端故障排除：当前访问服务器的主机
}
**/
