package filex

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"workbench-server/common/errorx"
	"workbench-server/common/global"
	"workbench-server/common/tool"
)

const MaxFileSize = 10 << 20 // 10 MB

/** SaveHttpImg 保存http上传图片，返回图片url
 * r *http.Request http请求
 * saveDir string 保存目录
 * prefix string 文件名前缀
 */
func SaveHttpImg(r *http.Request, saveDir, prefix string) (url string, err error) {
	r.ParseMultipartForm(MaxFileSize)
	file, header, err := r.FormFile("fileData")
	if err != nil {
		return "", errorx.New("讀取上傳文件失敗", err)
	}
	defer file.Close()
	_, filetype, err := tool.FileNameAndType(header.Filename)
	if err != nil {
		return "", err
	}
	ok, err := tool.ImgTypeVerify(filetype)
	if !ok {
		return "", err
	}
	ctx := r.Context()
	uid, err := tool.GetJwtUid(ctx)
	if err != nil {
		return "", err
	}
	filename := tool.GenerateImgName(fmt.Sprintf("%s_%d_", prefix, uid))
	name := fmt.Sprint(filename, ".", filetype)
	filepath := path.Join(saveDir, name)
	if tool.FileExist(filepath) {
		return "", errors.New("上傳異常，請重新試一試")
	}
	tempFile, err := os.Create(filepath)
	if err != nil {
		return "", errorx.InternalError("文件保存失敗", err)
	}
	defer tempFile.Close()
	io.Copy(tempFile, file)
	url, err = ImgHttpUrl(filepath)
	if err != nil {
		return "", errorx.InternalError("文件保存失敗", err)
	}
	return url, nil
}

//ImgUrl 传入图片路径，返回图片访问url
func ImgHttpUrl(filepath string) (string, error) {
	n := len(global.CONFIG.ServerConfig.RootPath)
	m := len(filepath)
	if n > m {
		return "", errors.New("ImgUrl方法出错")
	}
	if n == m {
		return "", nil
	}
	res := path.Join(global.CONFIG.StaticConfig.Domain, filepath[n:])
	return res, nil
}
