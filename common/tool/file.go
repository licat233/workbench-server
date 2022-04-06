package tool

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
	"workbench-server/common/errorx"
	"workbench-server/common/global"

	"go.uber.org/zap"
)

//ImgTypeVerify 图片类型检查
func ImgTypeVerify(fileType string) (bool, error) {
	ok, err := regexp.MatchString("^(jpg|png|jpeg|gif|ico|svg|webp|bmp)?$", fileType)
	if err != nil {
		global.ZAP.Error("文件类型校验失败:", zap.Error(err))
		return false, errorx.InternalError("ImgTypeVerify()方法出错文件类型校验失败", err)
	}
	if !ok {
		return false, errorx.New("不支持當前圖片格式，受支持的格式：jpg|png|jpeg|gif|ico|svg|webp|bmp")
	}
	return true, nil
}

//FilenameValid 文件名检查
func FilenameValidate(filename string) error {
	n := len(filename)
	if n == 0 {
		return errors.New("該文件沒有名字")
	}
	if n > 64 {
		return errors.New("文件名過長")
	}
	if ok, _ := regexp.MatchString("^.{1,64}\\.[A-Za-z0-9]{2,5}?$", filename); !ok {
		return errors.New("未知文件類型")
	}
	return nil
}

// FileSize 获取文件大小
func FileSize(filePath string) (int64, bool) {
	fi, err := os.Lstat(filePath)
	if err == nil {
		return fi.Size(), !fi.IsDir()
	}
	return fi.Size(), !os.IsNotExist(err)
}

// FileExist 判断文件是否存在
func FileExist(path string) bool {
	fi, err := os.Lstat(path)
	if err == nil {
		return !fi.IsDir()
	}
	return !os.IsNotExist(err)
}

// DirExist 判断文件夹是否存在
func DirExist(path string) bool {
	fi, err := os.Lstat(path)
	if err == nil {
		return fi.IsDir()
	}
	return !os.IsNotExist(err)
}

//FileNameAndType 获取文件名与类型，同时会检查文件名是否合法，或者是否存在
func FileNameAndType(filename string) (fileName string, fileType string, err error) {
	err = FilenameValidate(filename)
	if err != nil {
		return
	}
	index := strings.LastIndex(filename, ".")
	if index == -1 || index == len(filename)-1 {
		err = errors.New("文件类型不存在")
		return
	}
	fileName = filename[:index]
	fileType = filename[index+1:]
	return
}

//GenerateImgName 生成图片名字
func GenerateImgName(prefix string) string {
	return fmt.Sprintf("%s%d", prefix, time.Now().UnixNano())
}
