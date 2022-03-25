package filex

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
	"workbench-server/common/errorx"

	"github.com/tal-tech/go-zero/core/logx"
)

const maxFileSize = 10 << 20 // 10 MB

func SaveFile(r *http.Request, savePath string, isImg bool) (string, error) {
	r.ParseMultipartForm(maxFileSize)
	file, handler, err := r.FormFile("fileData")
	if err != nil {
		// fmt.Println(err)
		return "", errorx.NewServerError("讀取上傳文件失敗", err)
	}
	defer file.Close()

	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)

	//判断文件类型
	filename := handler.Filename
	if isImg {
		if ok, _ := regexp.MatchString("^[^\\.].+\\.(jpg|png|jpeg|gif|ico|svg|webp|bmp)?$", filename); !ok {
			return "", errors.New("請上傳圖片")
		}
	}
	//文件名校验
	if err := FilenameValid(filename); err != nil {
		return "", err
	}
	filepath := path.Join(savePath, filename)
	//判断文件是否存在
	if existFile, err := os.Stat(filepath); !os.IsNotExist(err) { //文件名相等
		if existFile.Size() == handler.Size { //大小相等
			//如果文件已存在，就不写入了
			return filepath, nil
			// return respx.NewDefaultResponse(nil, filepath, "文件已存在")
		}
		//文件重命名
		filename = NewFileName(savePath, filename)
	}

	filepath = path.Join(savePath, filename)
	// if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
	// 	logx.Error("upload接口创建文件夹失败：", err.Error())
	// 	return nil, errorx.NewServerError("文件保存失敗", err)
	// }

	tempFile, err := os.Create(filepath)
	if err != nil {
		logx.Error("upload接口创建文件失败：", err.Error())
		return "", errorx.NewServerError("文件保存失敗", err)
	}
	defer tempFile.Close()
	io.Copy(tempFile, file)
	return filepath, nil
}

/**
 * checkFileIsExist 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//FilenameValid 文件名校验
func FilenameValid(filename string) error {
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

//NewFileName 生成新的文件名
func NewFileName(savePath, filename string) string {
	index := strings.LastIndex(filename, ".")
	ftype := filename[index:]

	fname := filename[:index]
	num := 1
	newname := fmt.Sprint(fname, num, ftype)
	newpath := path.Join(savePath, newname)
	//检测是否存在，加入 i 作为限制器，防止意外陷入死循环
	i := 0
	for i < 100 && CheckFileIsExist(newpath) {
		newname = fmt.Sprint(fname, num+1, ftype)
		newpath = path.Join(savePath, newname)
		i++
	}
	if i == 100 && CheckFileIsExist(newpath) {
		nowTimestamp := int32(time.Now().Unix())
		newname = fmt.Sprint(fname, nowTimestamp, ftype)
	}
	return newname
}
