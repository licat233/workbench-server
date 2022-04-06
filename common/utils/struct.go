package utils

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"
	"time"
)

/** StructAssign 合并两个数据相同的不同结构体，将source的值转移到target
 * @target type interface 要修改的结构体指针
 * @source type interface 有数据的结构体指针
 */
func StructAssign(target interface{}, source interface{}) {
	targetVal := reflect.ValueOf(target).Elem() //获取target的值
	sourceVal := reflect.ValueOf(source).Elem() //获取source的值
	vTypeOfT := sourceVal.Type()                //获取source的结构体
	fieldNum := sourceVal.NumField()            //获取source的字段数量
	for i := 0; i < fieldNum; i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		sname := vTypeOfT.Field(i).Name       //获取source字段的名
		tname := targetVal.FieldByName(sname) //获取target字段的名
		if tname.IsValid() {                  // 验证target是否存在名称一模一样的字段
			snameType := vTypeOfT.Field(i).Type //获取source字段的类型
			tnameType := tname.Type()           //获取target字段的类型
			if tnameType == snameType {         // 验证target是否存在类型一模一样的字段
				tname.Set(sourceVal.Field(i)) //赋值给target的字段
			} else if snameType.Name() == "Time" && tnameType.Name() == "string" {
				toString := sourceVal.Field(i).Interface().(time.Time).Format("2006-01-02 15:04:05")
				tname.Set(reflect.ValueOf(toString))
			} else if snameType.Name() == "string" && tnameType.Name() == "Time" {
				if toTime, err := time.ParseInLocation("2006-01-02 15:04:05", sourceVal.Field(i).Interface().(string), time.Local); err == nil {
					tname.Set(reflect.ValueOf(toTime))
				}
			}
		}
	}
}

//StrctToSlice 结构体转切片
func StrctToSlice(f interface{}) []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		ss[i] = fmt.Sprintf("%v", v.Field(i))
	}
	return ss
}

/** PathHandler 上传路径处理
 * @config interface{} 路径配置结构体指针，{Path:string, any...}
 * @rootDir string 根目录
 */
func PathHandler(config any, rootDir string) {
	dataType := reflect.TypeOf(config)
	if dataType.Kind() != reflect.Ptr {
		//传入data不是指针类型
		panic("PathHandler：请传入指针类型的data")
	}
	if dataType.Elem().Kind() != reflect.Struct {
		//"data不是结构体类型
		return
	}
	dataVal := reflect.ValueOf(config).Elem()
	pathField := dataVal.FieldByName("Path")
	if !pathField.IsValid() {
		//path不字段不存在
		return
	}

	if pathField.Type().Kind() != reflect.String {
		//path不是string类型
		return
	}
	pathVal := pathField.String()
	dir := path.Join(rootDir, pathVal)
	pathField.SetString(dir) //整个方法的核心部分
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("创建文件夹【%s】失败:%v", dir, err))
	}

	n := dataVal.NumField()
	for i := 0; i < n; i++ {
		field := dataVal.Field(i)
		if field.Type().Kind() == reflect.Struct {
			//当前结构体Path字段值为空字符串，默认设为结构体名称
			childPathField := field.FieldByName("Path")
			if childPathField.IsValid() { //有path才处理
				if childPathField.Type().Kind() == reflect.String { //且为string格式
					childPathFieldVal := strings.TrimSpace(childPathField.String())
					if len(childPathFieldVal) == 0 { //path为空字符串，设为其结构体名称
						childrenPath := dataVal.Type().Field(i)
						childPathVal := strings.ToLower(childrenPath.Name)
						childPathField.SetString(childPathVal)
					}
					PathHandler(field.Addr().Interface(), dir)
				}
			}
		}
	}
}
