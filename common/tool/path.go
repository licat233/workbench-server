package tool

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

//获取当前项目地址
func CurrentProjectRootPath() string {
	fullPath := getCurrentAbPath()
	reg := regexp.MustCompile("^.*workbench-server")
	matched := reg.FindString(fullPath)
	return matched
}

//获取当前项目地址
func CurrentPath() string {
	return getCurrentAbPath()
}

// 最终方案-全兼容
func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	//如果是 go run 启动的
	tmpPath := getTmpDir()
	if tmpPath == "." {
		return getCurrentAbPathByCaller()
	}
	if strings.Contains(dir, getTmpDir()) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径（仅支持go build）
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（仅支持go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

/** StaticServerPathHandler 静态资源目录初始化
 * config type any 路径配置结构体指针，{Path:string, any...}
 * rootDir type string 根目录
 */
func StaticServerPathHandler(config any, rootDir string) {
	dataType := reflect.TypeOf(config)
	if dataType.Kind() != reflect.Ptr {
		//传入data不是指针类型
		panic("PathHandler：请传入config结构体的指针")
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
					StaticServerPathHandler(field.Addr().Interface(), dir)
				}
			}
		}
	}
}
