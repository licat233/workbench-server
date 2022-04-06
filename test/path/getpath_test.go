package main

import (
	"fmt"
	"testing"

	"github.com/licat233/goutil/getpath"
)

func TestGetPath(t *testing.T) {
	res := getpath.CurrentPath()
	fmt.Println(res)
}

// //CurrentProjectPath 返回当前项目地址
// func Test_CurrentProjectPath() (root string) {
// 	fullPath := getpath.CurrentPath()
// 	regexp.MatchString(/^/)
// 	return
// }
