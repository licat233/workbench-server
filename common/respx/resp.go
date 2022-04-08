package respx

import (
	"errors"
	"reflect"
	"workbench-server/common/errorx"
)

/** 与前端约定的数据规范
interface ErrorInfoStructure {
	success: boolean; // 请求状态
	data?: any; // 返回数据
	total?: number; // 数据总个数
	pageSize: number //单页数量
	current?: number; // 当前页码
	page?: number; // 【自己加上去的】总共有多少页，根据前端的pageSize来计算
}
**/

/**
* 前端要求的返回格式规范
  const result = {
   data: dataSource,
   total: tableListDataSource.length,
   success: true,
   pageSize: finalPageSize,
   current: parseInt(`${params.currentPage}`, 10) || 1,
 };
*/

/** Respx 对接前端antpro框架的接口格式规范
 */
// type Respx struct {
// 	Status   bool   `json:"success"`  // 【必需】请求状态
// 	Message  string `json:"message"`  // 【必需】响应消息
// 	Data     any    `json:"data"`     // 【必需】返回数据
// 	Total    int64  `json:"total"`    // 【选填】数据总个数
// 	PageSize int64  `json:"pageSize"` // 【选填】单页数量
// 	Current  int64  `json:"current"`  // 【选填】当前页码
// 	Page     int64  `json:"page"`     // 【自定】总共有多少页，根据前端的pageSize来计算
// }

var (
	SUCCESS_MSG = "请求成功"
)

//StateData 状态响应结构体
type StateData struct {
	Status  bool   `json:"success"` // 请求状态。
	Message string `json:"message"` // 响应消息
}

//SingleData 单数据响应结构体
type SingleData struct {
	Status  bool   `json:"success"` // 请求状态。可有可无
	Message string `json:"message"` // 响应消息
	Data    any    `json:"data"`    // 返回数据
}

//ListData 列表响应数据结构体
type ListData struct {
	Status   bool   `json:"success"`  // 【必需】请求状态
	Message  string `json:"message"`  // 【必需】响应消息
	Data     any    `json:"data"`     // 【选填】返回数据
	Total    int64  `json:"total"`    // 【选填】数据总个数
	PageSize int64  `json:"pageSize"` // 【选填】单页数量
	Current  int64  `json:"current"`  // 【选填】当前页码
	Page     int64  `json:"page"`     // 【自定】总共有多少页，根据前端的pageSize来计算
}

/** Result 想用语法糖，根据err响应信息，用于对接go-zero返回体
 * err error
 * params请依次传入：
 * alertMsg string 提示信息
 * data any 返回的数据，如果是分页请求，请返回slice类型
 * totalRecord in64 总数据量
 * pageSize int64 单页数据量
 * Current int64 当前页码
 */
func Result(err error, params ...any) (any, error) {
	// success := "请求成功"
	// fail := "请求失败"
	if err != nil {
		//返回错误
		return nil, err
	}
	return New(params), err
}

/** New 响应语法糖，新建成功响应
 * 该方法是无错误返回，所以不用传入err
 * 请依次传入：
 * alertMsg string 提示信息
 * data any 返回的数据，如果是分页请求，请返回slice类型
 * totalRecord in64 总数据量
 * pageSize int64 单页数据量
 * Current int64 当前页码
 * 如果不熟悉，请选择其它方法：StateResp、SingleResp、ListResp
 */
func New(params ...any) any {
	var data any
	alertMsg, totalRecord, pageSize, current := SUCCESS_MSG, int64(1), int64(1), int64(1)
	n := len(params)
	//参数只有1个，表示状态响应
	if n > 0 {
		if v, ok := params[0].(string); ok {
			if len(v) > 0 {
				alertMsg = v
			}
		} else {
			return errorx.InternalError("服务器错误", errors.New("respx.New()的第1个参数不是string"))
		}
	}
	if n == 0 || n == 1 {
		return StateResp(alertMsg)
	}
	//参数2个以上，表示单页数据
	if n > 1 {
		data = params[1]
	}
	if n == 2 {
		return SingleResp(alertMsg, data)
	}
	//参数3个以上，表示分页数据
	if n > 2 {
		if v, ok := params[2].(int64); ok {
			totalRecord = v
		} else {
			return errorx.InternalError("服务器错误", errors.New("respx.New()的第3个参数totalRecord不是int64"))
		}
	}
	if n > 3 {
		if v, ok := params[3].(int64); ok {
			pageSize = v
		} else {
			return errorx.InternalError("服务器错误", errors.New("respx.New()的第4个参数pageSize不是int64"))
		}
	}
	if n > 4 {
		if v, ok := params[4].(int64); ok {
			current = v
		} else {
			return errorx.InternalError("服务器错误", errors.New("respx.New()的第5个参数current不是int64"))
		}
	}
	return ListResp(alertMsg, data, totalRecord, pageSize, current)
}

//StateResp 返回响应状态
func StateResp(alertMsg string) *StateData {
	return &StateData{
		Status:  true,
		Message: alertMsg,
	}
}

//StateResp 默认返回响应状态，主要用于对接go-zero
func DefaultStateResp(err error) (*StateData, error) {
	if err != nil {
		return nil, err
	}
	return &StateData{
		Status:  true,
		Message: SUCCESS_MSG,
	}, nil
}

//SingleResp 返回单数据响应
func SingleResp(alertMsg string, data any) *SingleData {
	return &SingleData{
		Status:  true,
		Message: alertMsg,
		Data:    data,
	}
}

//SingleResp 默认返回单数据响应，主要用于对接go-zero
func DefaultSingleResp(err error, data any) (*SingleData, error) {
	if err != nil {
		return nil, err
	}
	return &SingleData{
		Status:  true,
		Message: SUCCESS_MSG,
		Data:    data,
	}, nil
}

/** ListResp 返回列表数据响应，用于分页查询数据
* alertMsg string 提示信息
* data any 返回数据
* totalRecord in64 总数据量
* pageSize int64 单页数据量
* Current int64 当前页码
 */
func ListResp(alertMsg string, listData any, totalRecord, pageSize, Current int64) *ListData {
	data := listData
	if listData != nil {
		if reflect.TypeOf(listData).Kind().String() != "slice" {
			data = []any{listData}
		}
	} else {
		data = []any{}
	}
	if pageSize == 0 {
		pageSize = 1
	}
	return &ListData{
		Status:   true,
		Message:  alertMsg,
		Data:     data,
		Total:    totalRecord,
		PageSize: pageSize,
		Current:  Current,
		Page:     (totalRecord + pageSize - 1) / pageSize,
	}
}

/** DefaultListResp 默认返回列表数据响应，用于分页查询数据，主要用于对接go-zero
* err error 提示信息
* data any 返回数据
* totalRecord in64 总数据量
* pageSize int64 单页数据量
* Current int64 当前页码
 */
func DefaultListResp(err error, listData any, totalRecord, pageSize, Current int64) (*ListData, error) {
	if err != nil {
		return nil, err
	}
	data := listData
	if listData != nil {
		if reflect.TypeOf(listData).Kind().String() != "slice" {
			data = []any{listData}
		}
	} else {
		data = []any{}
	}
	if pageSize == 0 {
		pageSize = 1
	}
	return &ListData{
		Status:   true,
		Message:  SUCCESS_MSG,
		Data:     data,
		Total:    totalRecord,
		PageSize: pageSize,
		Current:  Current,
		Page:     (totalRecord + pageSize - 1) / pageSize,
	}, nil
}

//Msg 根据error返回不同提示信息
func Msg(err error, data any, successMsg, failedMsg string) (any, error) {
	if err != nil {
		return nil, errorx.InternalError(failedMsg, err)
	}
	return SingleResp(successMsg, data), nil
}
