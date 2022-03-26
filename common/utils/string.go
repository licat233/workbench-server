package utils

import (
	"fmt"
	"strconv"
	"strings"
)

//StringSplice 字符串拼接
func StringSplice(strs ...string) string {
	var build strings.Builder
	for _, str := range strs {
		build.WriteString(str)
	}
	return build.String()
}

//Int64ToStr int64转string
func Int64ToStr(id int64) (str string) {
	str = fmt.Sprintf("%d", id)
	return
}

//StrToInt64 字符串转int64
func StrToInt64(str string) (id int64) {
	id, _ = strconv.ParseInt(str, 10, 64)
	return
}

